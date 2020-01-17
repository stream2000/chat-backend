package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _MembershipMgr struct {
	*_BaseMgr
}

// MembershipMgr open func
func MembershipMgr(db *gorm.DB) *_MembershipMgr {
	if db == nil {
		panic(fmt.Errorf("MembershipMgr need init by db"))
	}
	return &_MembershipMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MembershipMgr) GetTableName() string {
	return "membership"
}

// Get 获取
func (obj *_MembershipMgr) Get() (result Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MembershipMgr) Gets() (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithGroupID group_id获取
func (obj *_MembershipMgr) WithGroupID(GroupID string) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = GroupID })
}

// WithUserID user_id获取
func (obj *_MembershipMgr) WithUserID(UserID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = UserID })
}

// WithAdminID admin_id获取 允许进群的管理员的id
func (obj *_MembershipMgr) WithAdminID(AdminID string) Option {
	return optionFunc(func(o *options) { o.query["admin_id"] = AdminID })
}

// WithCreatedAt created_at获取
func (obj *_MembershipMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_MembershipMgr) GetByOption(opts ...Option) (result Membership, err error) {
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
func (obj *_MembershipMgr) GetByOptions(opts ...Option) (results []*Membership, err error) {
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

// GetFromGroupID 通过group_id获取内容
func (obj *_MembershipMgr) GetFromGroupID(GroupID string) (result Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", GroupID).Find(&result).Error

	return
}

// GetBatchFromGroupID 批量唯一主键查找
func (obj *_MembershipMgr) GetBatchFromGroupID(GroupIDs []string) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id IN (?)", GroupIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_MembershipMgr) GetFromUserID(UserID string) (result Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", UserID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_MembershipMgr) GetBatchFromUserID(UserIDs []string) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", UserIDs).Find(&results).Error

	return
}

// GetFromAdminID 通过admin_id获取内容 允许进群的管理员的id
func (obj *_MembershipMgr) GetFromAdminID(AdminID string) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("admin_id = ?", AdminID).Find(&results).Error

	return
}

// GetBatchFromAdminID 批量唯一主键查找 允许进群的管理员的id
func (obj *_MembershipMgr) GetBatchFromAdminID(AdminIDs []string) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("admin_id IN (?)", AdminIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MembershipMgr) GetFromCreatedAt(CreatedAt time.Time) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_MembershipMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MembershipMgr) FetchByPrimaryKey(GroupID string, UserID string) (result Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ? AND user_id = ?", GroupID, UserID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_MembershipMgr) FetchByIndex(AdminID string) (results []*Membership, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("admin_id = ?", AdminID).Find(&results).Error

	return
}
