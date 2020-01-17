package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _TagMgr struct {
	*_BaseMgr
}

// TagMgr open func
func TagMgr(db *gorm.DB) *_TagMgr {
	if db == nil {
		panic(fmt.Errorf("TagMgr need init by db"))
	}
	return &_TagMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagMgr) GetTableName() string {
	return "tag"
}

// Get 获取
func (obj *_TagMgr) Get() (result Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagMgr) Gets() (results []*Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TagMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithName name获取
func (obj *_TagMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithCreatedAt created_at获取
func (obj *_TagMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_TagMgr) GetByOption(opts ...Option) (result Tag, err error) {
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
func (obj *_TagMgr) GetByOptions(opts ...Option) (results []*Tag, err error) {
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
func (obj *_TagMgr) GetFromID(ID int) (result Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromID(IDs []int) (results []*Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TagMgr) GetFromName(Name string) (result Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&result).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromName(Names []string) (results []*Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_TagMgr) GetFromCreatedAt(CreatedAt time.Time) (results []*Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TagMgr) FetchByPrimaryKey(ID int) (result Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_TagMgr) FetchByUnique(Name string) (result Tag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&result).Error

	return
}
