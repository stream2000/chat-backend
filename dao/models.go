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
	Id        int
	Account   string
	PassWd    string `gorm:"column:passwd"`
	AvatarUrl string
}

type Consume struct {
	SenderId   int
	ReceiverId int
	Offset     int
}
