/*
@Time : 2020/6/11 20:51
@Author : Minus4
*/
package websocket

import (
	"github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
	"m4-im/pkg/util"
	"strconv"
	"strings"
)

func badCondition() {
	panic("this should not happen. So I use panic here. And I may not going to fix it in the future")
}

func OnAuth(s socketio.Conn, param AuthParam) bool {
	claims, err := util.ParseToken(param.Jwt)
	if err != nil {
		return false
	}
	c := &channel{
		socket:    s,
		closeChan: make(chan struct{}, 1),
	}
	log.Info("User with id ", claims.Id, " signed in")
	channelManager.addNewChannel(c, claims.Id)
	go c.serve()
	return true
}

func OnLogOff(s socketio.Conn) {
	if c, ok := channelManager.getChannelByCid(s.ID()); ok {
		log.Info("User with id ", c.u.id, " signed off")
		channelManager.removeChannelByCid(s.ID())
	} else {
		log.Fatal("can't get channel by cid ", s.ID())
	}
}

func OnDisconnect(s socketio.Conn, msg string) {
	if c, ok := channelManager.getChannelByCid(s.ID()); ok {
		channelManager.removeChannelByCid(s.ID())
		log.Info("User with id ", c.u.id, " disconnect")
	}
}

func OnDefaultGroupMessage(s socketio.Conn, msg PushMessage) {
	msg.Text = strings.Replace(msg.Text, "\n", "", -1)
	senderChan, _ := channelManager.getChannelByCid(s.ID())
	channelManager.Lock()
	for _, u := range channelManager.users {
		if u.id == senderChan.u.id {
			continue
		}
		u.sendNewGroupMessage(senderChan.u.id, msg.Text)
	}
	channelManager.Unlock()
}

func OnNewMessage(s socketio.Conn, msg PushMessage) {
	msg.Text = strings.Replace(msg.Text, "\n", "", -1)
	if receiver, ok := channelManager.getUserByUserId(strconv.Itoa(msg.ReceiverId)); ok {
		senderChan, _ := channelManager.getChannelByCid(s.ID())
		log.Infof("[%s:%s] -> [%s:%s] : %s", senderChan.u.name, senderChan.u.id, receiver.name, receiver.id, msg.Text)
		receiver.sendNewMessage(senderChan.u.id, msg.Text)
	} else {
		badCondition()
	}
}
