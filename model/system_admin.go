package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _SystemAdminMgr struct {
	*_BaseMgr
}

// SystemAdminMgr open func
func SystemAdminMgr(db *gorm.DB) *_SystemAdminMgr {
	if db == nil {
		panic(fmt.Errorf("SystemAdminMgr need init by db"))
	}
	return &_SystemAdminMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemAdminMgr) GetTableName() string {
	return "system_admin"
}

// Get 获取
func (obj *_SystemAdminMgr) Get() (result SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemAdminMgr) Gets() (results []*SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_SystemAdminMgr) WithUserID(UserID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithRoleID role_id获取
func (obj *_SystemAdminMgr) WithRoleID(RoleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = RoleID })
}

// GetByOption 功能选项模式获取
func (obj *_SystemAdminMgr) GetByOption(opts ...Option) (result SystemAdmin, err error) {
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
func (obj *_SystemAdminMgr) GetByOptions(opts ...Option) (results []*SystemAdmin, err error) {
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
func (obj *_SystemAdminMgr) GetFromUserID(UserID string) (result SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_SystemAdminMgr) GetBatchFromUserID(UserIDs []string) (results []*SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_SystemAdminMgr) GetFromRoleID(RoleID int) (result SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id = ?", RoleID).Find(&result).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找
func (obj *_SystemAdminMgr) GetBatchFromRoleID(RoleIDs []int) (results []*SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("role_id IN (?)", RoleIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemAdminMgr) FetchByPrimaryKey(UserID string, RoleID int) (result SystemAdmin, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ? AND role_id = ?", UserID, RoleID).Find(&result).Error

	return
}
