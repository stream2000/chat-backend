/*
@Time : 2020/6/11 20:51
@Author : Minus4
*/
package websocket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"m4-im/pkg/util"
)

func OnAuth(s socketio.Conn, param AuthParam) bool {
	claims, err := util.ParseToken(param.Jwt)
	if err != nil {
		return false
	}
	fmt.Printf("socket id : %s client id : %s\n", s.ID(), claims.Id)
	return true
}

type LogOffParam struct {
	Msg string `json:"msg"`
}

func OnLogOff(s socketio.Conn) {
	fmt.Printf("client with id %s log off\n", s.ID())
}
