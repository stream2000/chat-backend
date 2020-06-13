/*
@Time : 2020/6/11 20:59
@Author : Minus4
*/
package websocket

import (
	"github.com/googollee/go-socket.io"
	"github.com/sirupsen/logrus"
	"m4-im/dao"
	"strconv"
)

type user struct {
	id           string
	name         string
	messageQueue chan message
}

func (u user) sendNewMessage(sender string, text string) {
	m := message{
		topic: "msg",
		body: map[string]interface{}{
			"sender_id": sender,
			"text":      text,
		},
	}
	select {
	case u.messageQueue <- m:
	default:
		logrus.Fatal("message is discard because the queue is full")
	}
}

type message struct {
	topic string
	body  interface{}
}

type channel struct {
	userId         string
	receiveChannel chan message
	socket         socketio.Conn
	closeChan      chan struct{}
}

func (c *channel) close() {
	logrus.Info("Channel with id ", c.userId, " is closed")
	select {
	case c.closeChan <- struct{}{}:
	default:
	}
}

func (c *channel) serve() {
	logrus.Info("Channel with id ", c.userId, " is serving")
	for {
		select {
		case <-c.closeChan:
			return
		case m := <-c.receiveChannel:
			logrus.Debugf("send message [%s] ", m.body)
			c.socket.Emit(m.topic, m.body)
		}
	}
}

func (c *channel) sendNewMessage(text string) {
	m := message{
		topic: "msg",
		body:  text,
	}
	select {
	case c.receiveChannel <- m:
	default:
		logrus.Fatal("message is discard because the queue is full")
	}
}

func (c *channel) notifyNewUser(u dao.User) {
	m := message{
		topic: "newUser",
		body: map[string]interface{}{
			"id":   strconv.Itoa(u.Id),
			"img":  u.AvatarUrl,
			"name": u.AvatarUrl,
		},
	}
	select {
	case c.receiveChannel <- m:
	default:
		logrus.Fatalf("Message [%+v]  is discard because the queue is full", m)
	}
}

func (c *channel) notifyUserOnline(id int) {
	m := message{
		topic: "online",
		body:  id,
	}
	select {
	case c.receiveChannel <- m:
	default:
		logrus.Fatalf("Message [%+v]  is discard because the queue is full", m)
	}
}

func (c *channel) notifyUserOffline(id int) {
	m := message{
		topic: "offline",
		body:  id,
	}
	select {
	case c.receiveChannel <- m:
	default:
		logrus.Fatalf("Message [%+v]  is discard because the queue is full", m)
	}
}
