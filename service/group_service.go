/*
@Time : 2020/1/18 01:10
@Author : Minus4
*/
package service

import (
	"github.com/jinzhu/gorm"
	"m4-im/dao"
	"m4-im/model"
	"m4-im/pkg/e"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
)

func CreatGroup(param params.AddNewGroupParam, userId string) (err error) {
	db := dao.GetDB()
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	// #1 create group
	group := &model.ChatGroup{
		ID:        util.UUID(),
		GroupName: param.GroupName,
		CreatorID: userId,
		Info:      param.Info,
		Tags:      "default",
		Status:    StatusOK,
	}
	err = tx.Create(group).Error
	if err != nil {
		return NewServiceError(e.ErrUnKnownInternalError, err)
	}

	// #2 add admin privilege
	admin := &model.GroupAdmin{
		UserID:   userId,
		GroupID:  group.ID,
		RoleCode: GroupAdminUser,
	}
	err = tx.Create(admin).Error
	if err != nil {
		return NewServiceError(e.ErrUnKnownInternalError, err)
	}

	return
}

func ExistsGroupById(id string) bool {
	if _, err := GetGroupById(id); err != nil {
		return false
	} else {
		return true
	}
}

func DeleteGroup(id string) (err error) {
	group, err := GetGroupById(id)
	if err != nil {
		return err
	}
	group.Status = StatusDeleted
	err = dao.GetDB().Table("chat_group").Save(&group).Error
	if err != nil {
		return NewServiceError(e.ErrUnKnownInternalError, err)
	}
	return
}

func GetGroupById(id string) (model.ChatGroup, error) {
	groupRepo := model.ChatGroupMgr(dao.GetDB())
	group, err := groupRepo.GetFromID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return group, NewServiceError(e.ErrGroupNotFound, err)
		} else {
			return group, NewServiceError(e.ErrUnKnownInternalError, err)
		}
	}
	if group.Status == StatusDeleted {
		return group, NewServiceError(e.ErrGroupNotFound, err)
	}
	return group, nil
}
