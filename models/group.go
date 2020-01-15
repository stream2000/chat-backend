package models

import (
	"fmt"
)

type _GroupMgr struct {
	*BaseMgr
}

// GroupMgr open func
func GroupMgr() *_GroupMgr {
	if db == nil {
		panic(fmt.Errorf("GroupMgr need init by db"))
	}
	return &_GroupMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupMgr) GetTableName() string {
	return "group"
}

// Get 获取
func (obj *_GroupMgr) Get() (result Group, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupMgr) Gets() (results []*Group, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// GetByOption 功能选项模式获取
func (obj *_GroupMgr) GetByOption(opts ...Option) (result Group, err error) {
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
func (obj *_GroupMgr) GetByOptions(opts ...Option) (results []*Group, err error) {
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

//////////////////////////primary index case ////////////////////////////////////////////
