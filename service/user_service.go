/*
@Time : 2020/1/15 21:54
@Author : Minus4
*/
package service

import (
	"m4-im/dao"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
)

func Register(param params.RegisterParam) (err error) {
	err = dao.AddUser(param)
	return
}

func Auth(email, password string) (bool, string) {
	userInfo, err := dao.GetUserInfoByEmail(email)
	if err != nil {
		return false, ""
	}
	storedSum := userInfo.Password
	currentSum := util.EncryptAccount(email, password)

	return storedSum == currentSum, userInfo.ID
}
