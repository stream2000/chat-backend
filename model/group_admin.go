package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _GroupAdminMgr struct {
	*_BaseMgr
}

// GroupAdminMgr open func
func GroupAdminMgr(db *gorm.DB) *_GroupAdminMgr {
	if db == nil {
		panic(fmt.Errorf("GroupAdminMgr need init by db"))
	}
	return &_GroupAdminMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupAdminMgr) GetTableName() string {
	return "group_admin"
}

// Get 获取
func (obj *_GroupAdminMgr) Get() (result GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupAdminMgr) Gets() (results []*GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_GroupAdminMgr) WithUserID(UserID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithGroupID group_id获取
func (obj *_GroupAdminMgr) WithGroupID(GroupID string) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = GroupID })
}

// WithRoleCode role_code获取
func (obj *_GroupAdminMgr) WithRoleCode(RoleCode int) Option {
	return optionFunc(func(o *options) { o.query["role_code"] = RoleCode })
}

// GetByOption 功能选项模式获取
func (obj *_GroupAdminMgr) GetByOption(opts ...Option) (result GroupAdmin, err error) {
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
func (obj *_GroupAdminMgr) GetByOptions(opts ...Option) (results []*GroupAdmin, err error) {
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
func (obj *_GroupAdminMgr) GetFromUserID(UserID string) (result GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_GroupAdminMgr) GetBatchFromUserID(UserIDs []string) (results []*GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_GroupAdminMgr) GetFromGroupID(GroupID string) (result GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", GroupID).Find(&result).Error

	return
}

// GetBatchFromGroupID 批量唯一主键查找
func (obj *_GroupAdminMgr) GetBatchFromGroupID(GroupIDs []string) (results []*GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id IN (?)", GroupIDs).Find(&results).Error

	return
}

// GetFromRoleCode 通过role_code获取内容
func (obj *_GroupAdminMgr) GetFromRoleCode(RoleCode int) (results []*GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_code = ?", RoleCode).Find(&results).Error

	return
}

// GetBatchFromRoleCode 批量唯一主键查找
func (obj *_GroupAdminMgr) GetBatchFromRoleCode(RoleCodes []int) (results []*GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_code IN (?)", RoleCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupAdminMgr) FetchByPrimaryKey(UserID string, GroupID string) (result GroupAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ? AND group_id = ?", UserID, GroupID).Find(&result).Error

	return
}
