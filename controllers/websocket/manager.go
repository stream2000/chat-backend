/*
@Time : 2020/6/11 20:47
@Author : Minus4
*/
package websocket

import (
	"github.com/sirupsen/logrus"
	"m4-im/dao"
	"strconv"
	"sync"
)

var channelManager = &ChannelManager{
	cIdToChannel:    make(map[string]*channel),
	userIdToChannel: make(map[string]*channel),
	users:           map[string]user{},
}

func SetUpChannelManager() {
	db := dao.GetDB()
	dbUsers := make([]dao.User, 0)
	db.Table("user").Find(&dbUsers)
	for _, u := range dbUsers {
		AddNewUser(u)
	}
}

func AddNewUser(u dao.User) {
	user := user{
		id:           strconv.Itoa(u.Id),
		messageQueue: make(chan message, 10000),
		name:         u.Account,
	}
	channelManager.Lock()
	for _, currentUser := range channelManager.userIdToChannel {
		currentUser.notifyNewUser(u)
	}
	channelManager.users[strconv.Itoa(u.Id)] = user
	channelManager.Unlock()
}

type ChannelManager struct {
	sync.Mutex
	users           map[string]user
	cIdToChannel    map[string]*channel
	userIdToChannel map[string]*channel
}

func (m *ChannelManager) addNewChannel(c *channel) {
	logrus.Info("Add new channel with id ", c.userId)
	m.Lock()
	if u, ok := m.users[c.userId]; ok {
		c.receiveChannel = u.messageQueue
	} else {
		panic("unknown new user !!!!!")
	}
	for _, c := range channelManager.userIdToChannel {
		id, _ := strconv.ParseInt(c.userId, 10, 32)
		c.notifyUserOffline(int(id))
	}
	m.cIdToChannel[c.socket.ID()] = c
	m.userIdToChannel[c.userId] = c
	m.Unlock()
}

func (m *ChannelManager) removeChannelByCid(id string) {
	logrus.Info("Remove channel with user id: ", id)
	m.Lock()
	defer m.Unlock()
	if channel, ok := m.cIdToChannel[id]; ok {
		delete(m.userIdToChannel, id)
		delete(m.cIdToChannel, channel.socket.ID())
		channel.close()
		idInInt, _ := strconv.ParseInt(id, 10, 32)
		for _, c := range m.userIdToChannel {
			c.notifyUserOffline(int(idInInt))
		}
	}
}

func (m *ChannelManager) getChannelByUserId(id string) (*channel, bool) {
	m.Lock()
	defer m.Unlock()
	c, ok := m.userIdToChannel[id]
	return c, ok
}

func (m *ChannelManager) getChannelByCid(id string) (*channel, bool) {
	m.Lock()
	defer m.Unlock()
	c, ok := m.cIdToChannel[id]
	return c, ok
}

func (m *ChannelManager) getUserByUserId(id string) (user, bool) {
	m.Lock()
	defer m.Unlock()
	u, ok := m.users[id]
	return u, ok
}
