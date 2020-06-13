/*
@Time : 2020/6/11 20:50
@Author : Minus4
*/
package websocket

type AuthParam struct {
	Jwt string `json:"jwt"`
}

type LogOffParam struct {
	Msg string `json:"msg"`
}

type PushMessage struct {
	Text       string `json:"text"`
	ReceiverId int    `json:"receiver_id"`
}
