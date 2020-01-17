package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _SysRoleMgr struct {
	*_BaseMgr
}

// SysRoleMgr open func
func SysRoleMgr(db *gorm.DB) *_SysRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleMgr need init by db"))
	}
	return &_SysRoleMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleMgr) GetTableName() string {
	return "sys_role"
}

// Get 获取
func (obj *_SysRoleMgr) Get() (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleMgr) Gets() (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SysRoleMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithName name获取
func (obj *_SysRoleMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithCreatorID creator_id获取
func (obj *_SysRoleMgr) WithCreatorID(CreatorID string) Option {
	return optionFunc(func(o *options) { o.query["creator_id"] = CreatorID })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleMgr) GetByOption(opts ...Option) (result SysRole, err error) {
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
func (obj *_SysRoleMgr) GetByOptions(opts ...Option) (results []*SysRole, err error) {
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

// GetFromID 通过id获取内容
func (obj *_SysRoleMgr) GetFromID(ID int) (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromID(IDs []int) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_SysRoleMgr) GetFromName(Name string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromName(Names []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromCreatorID 通过creator_id获取内容
func (obj *_SysRoleMgr) GetFromCreatorID(CreatorID string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id = ?", CreatorID).Find(&results).Error

	return
}

// GetBatchFromCreatorID 批量唯一主键查找
func (obj *_SysRoleMgr) GetBatchFromCreatorID(CreatorIDs []string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id IN (?)", CreatorIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleMgr) FetchByPrimaryKey(ID int) (result SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_SysRoleMgr) FetchByIndex(CreatorID string) (results []*SysRole, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id = ?", CreatorID).Find(&results).Error

	return
}
