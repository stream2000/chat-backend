package models

import (
	"fmt"
)

type _AdminMgr struct {
	*BaseMgr
}

// AdminMgr open func
func AdminMgr() *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	return &_AdminMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "admin"
}

// Get 获取
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_AdminMgr) WithUserID(UserID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithGroupID group_id获取
func (obj *_AdminMgr) WithGroupID(GroupID string) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = GroupID })
}

// WithRoleCode role_code获取
func (obj *_AdminMgr) WithRoleCode(RoleCode int) Option {
	return optionFunc(func(o *options) { o.query["role_code"] = RoleCode })
}

// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过user_id获取内容
func (obj *_AdminMgr) GetFromUserID(UserID string) (result Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_AdminMgr) GetBatchFromUserID(UserIDs []string) (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_AdminMgr) GetFromGroupID(GroupID string) (result Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", GroupID).Find(&result).Error

	return
}

// GetBatchFromGroupID 批量唯一主键查找
func (obj *_AdminMgr) GetBatchFromGroupID(GroupIDs []string) (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id IN (?)", GroupIDs).Find(&results).Error

	return
}

// GetFromRoleCode 通过role_code获取内容
func (obj *_AdminMgr) GetFromRoleCode(RoleCode int) (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_code = ?", RoleCode).Find(&results).Error

	return
}

// GetBatchFromRoleCode 批量唯一主键查找
func (obj *_AdminMgr) GetBatchFromRoleCode(RoleCodes []int) (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_code IN (?)", RoleCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByAdminPkUniqueIndex primay or index 获取唯一内容
func (obj *_AdminMgr) FetchByAdminPkUniqueIndex(UserID string, GroupID string) (result Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ? AND group_id = ?", UserID, GroupID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_AdminMgr) FetchByIndex(GroupID string) (results []*Admin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", GroupID).Find(&results).Error

	return
}
