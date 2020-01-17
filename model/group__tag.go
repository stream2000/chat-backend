package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _GroupTagMgr struct {
	*_BaseMgr
}

// GroupTagMgr open func
func GroupTagMgr(db *gorm.DB) *_GroupTagMgr {
	if db == nil {
		panic(fmt.Errorf("GroupTagMgr need init by db"))
	}
	return &_GroupTagMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupTagMgr) GetTableName() string {
	return "group__tag"
}

// Get 获取
func (obj *_GroupTagMgr) Get() (result GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupTagMgr) Gets() (results []*GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithGroupID group_id获取
func (obj *_GroupTagMgr) WithGroupID(GroupID string) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = GroupID })
}

// WithTagID tag_id获取
func (obj *_GroupTagMgr) WithTagID(TagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = TagID })
}

// GetByOption 功能选项模式获取
func (obj *_GroupTagMgr) GetByOption(opts ...Option) (result GroupTag, err error) {
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
func (obj *_GroupTagMgr) GetByOptions(opts ...Option) (results []*GroupTag, err error) {
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
func (obj *_GroupTagMgr) GetFromGroupID(GroupID string) (result GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ?", GroupID).Find(&result).Error

	return
}

// GetBatchFromGroupID 批量唯一主键查找
func (obj *_GroupTagMgr) GetBatchFromGroupID(GroupIDs []string) (results []*GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id IN (?)", GroupIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容
func (obj *_GroupTagMgr) GetFromTagID(TagID int) (result GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tag_id = ?", TagID).Find(&result).Error

	return
}

// GetBatchFromTagID 批量唯一主键查找
func (obj *_GroupTagMgr) GetBatchFromTagID(TagIDs []int) (results []*GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tag_id IN (?)", TagIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupTagMgr) FetchByPrimaryKey(GroupID string, TagID int) (result GroupTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_id = ? AND tag_id = ?", GroupID, TagID).Find(&result).Error

	return
}
