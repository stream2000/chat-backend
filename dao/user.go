/*
@Time : 2020/1/16 12:22
@Author : Minus4
*/
package dao

import (
	"fmt"
	"m4-im/model"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
)

func GetUserInfoByEmail(email string) (result model.User, err error) {
	userRepo := model.UserMgr(db)
	return userRepo.GetFromEmail(email)
}

func ExistsUserByEmail(email string) bool {
	userRepo := model.UserMgr(db)
	if _, err := userRepo.GetFromEmail(email); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func AddUser(param params.RegisterParam) (err error) {
	if ExistsUserByEmail(param.Email) {
		return fmt.Errorf("user with email %s already exists ", param.Email)
	}
	user := new(model.User)
	user.Email = param.Email
	user.Password = util.EncryptAccount(param.Email, param.Password)
	user.ID = util.UUID()
	user.Nickname = "user-" + user.ID
	db.Create(user)
	if !ExistsUserByEmail(param.Email) {
		return fmt.Errorf("error when creating user with email %s ", param.Email)
	}
	return
}
