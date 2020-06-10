/*
@Time : 2020/6/9 17:42
@Author : Minus4
*/
package dao

type Message struct {
	SenderId   int
	ReceiverId int
	Offset     int
	Text       string
}

type User struct {
	Id        int    `json:"id"`
	Account   string `json:"name"`
	PassWd    string `gorm:"column:passwd"`
	AvatarUrl string `json:"img"`
}

type Consume struct {
	SenderId   int
	ReceiverId int
	Offset     int
}
