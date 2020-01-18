/*
@Time : 2020/1/15 21:54
@Author : Minus4
*/
package service

import (
	"fmt"
	"m4-im/dao"
	"m4-im/model"
	"m4-im/pkg/e"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
	"time"
)

func Register(param params.RegisterParam) (err error) {
	if existsUserByEmail(param.Email) {
		return NewServiceError(e.ErrUserNotFound, err)
	}
	id := util.UUID()
	user := &model.User{
		ID:        id,
		Email:     param.Email,
		Password:  util.EncryptAccount(param.Email, param.Password),
		Nickname:  "user-" + id,
		CreatedAt: time.Time{},
		Status:    StatusOK,
	}
	err = dao.GetDB().Create(user).Error
	if err != nil {
		return NewServiceError(e.ErrUnKnownInternalError, err)
	}
	return
}

func Auth(email, password string) (bool, string) {
	userInfo, err := getUserInfoByEmail(email)
	if err != nil {
		return false, ""
	}
	storedSum := userInfo.Password
	currentSum := util.EncryptAccount(email, password)

	return storedSum == currentSum, userInfo.ID
}

func existsUserByEmail(email string) bool {
	var userRepo = model.UserMgr(dao.GetDB())
	if user, err := userRepo.GetFromEmail(email); err != nil {
		return false
	} else {
		// return false is user is deleted
		return user.Status == StatusOK
	}
}

func getUserInfoByEmail(email string) (result model.User, err error) {
	var userRepo = model.UserMgr(dao.GetDB())
	return userRepo.GetFromEmail(email)
}

func GetUserInfoById(id string) (result model.User, err error) {
	var userRepo = model.UserMgr(dao.GetDB())
	user, err := userRepo.GetFromID(id)

	if err != nil || user.Status != StatusOK {
		return user, NewServiceError(e.ErrUserNotFound, err)
	}
	return user, nil
}

func CheckUserRole(userId, groupId string, roleCode int) error {
	adminRepo := model.GroupAdminMgr(dao.GetDB())
	adminInfo, err := adminRepo.FetchByPrimaryKey(userId, groupId)
	if err != nil {
		return NewServiceError(e.ErrUnAuthorized, fmt.Errorf("用户不是小组的管理员！"))
	}
	if adminInfo.RoleCode == GroupManagerUser && roleCode == GroupAdminUser {
		return NewServiceError(e.ErrUnAuthorized, fmt.Errorf("用户不是群主"))
	}
	return nil
}
