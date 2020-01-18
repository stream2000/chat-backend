package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _ChatGroupMgr struct {
	*_BaseMgr
}

// ChatGroupMgr open func
func ChatGroupMgr(db *gorm.DB) *_ChatGroupMgr {
	if db == nil {
		panic(fmt.Errorf("ChatGroupMgr need init by db"))
	}
	return &_ChatGroupMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ChatGroupMgr) GetTableName() string {
	return "chat_group"
}

// Get 获取
func (obj *_ChatGroupMgr) Get() (result ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ChatGroupMgr) Gets() (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ChatGroupMgr) WithID(ID string) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithGroupName group_name获取
func (obj *_ChatGroupMgr) WithGroupName(GroupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = GroupName })
}

// WithCreatedAt created_at获取
func (obj *_ChatGroupMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// WithCreatorID creator_id获取
func (obj *_ChatGroupMgr) WithCreatorID(CreatorID string) Option {
	return optionFunc(func(o *options) { o.query["creator_id"] = CreatorID })
}

// WithInfo info获取
func (obj *_ChatGroupMgr) WithInfo(Info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = Info })
}

// WithTags tags获取
func (obj *_ChatGroupMgr) WithTags(Tags string) Option {
	return optionFunc(func(o *options) { o.query["tags"] = Tags })
}

// WithStatus status获取
func (obj *_ChatGroupMgr) WithStatus(Status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = Status })
}

// GetByOption 功能选项模式获取
func (obj *_ChatGroupMgr) GetByOption(opts ...Option) (result ChatGroup, err error) {
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
func (obj *_ChatGroupMgr) GetByOptions(opts ...Option) (results []*ChatGroup, err error) {
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
func (obj *_ChatGroupMgr) GetFromID(ID string) (result ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromID(IDs []string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容
func (obj *_ChatGroupMgr) GetFromGroupName(GroupName string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_name = ?", GroupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromGroupName(GroupNames []string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("group_name IN (?)", GroupNames).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ChatGroupMgr) GetFromCreatedAt(CreatedAt time.Time) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

// GetFromCreatorID 通过creator_id获取内容
func (obj *_ChatGroupMgr) GetFromCreatorID(CreatorID string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id = ?", CreatorID).Find(&results).Error

	return
}

// GetBatchFromCreatorID 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromCreatorID(CreatorIDs []string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id IN (?)", CreatorIDs).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_ChatGroupMgr) GetFromInfo(Info string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", Info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromInfo(Infos []string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", Infos).Find(&results).Error

	return
}

// GetFromTags 通过tags获取内容
func (obj *_ChatGroupMgr) GetFromTags(Tags string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tags = ?", Tags).Find(&results).Error

	return
}

// GetBatchFromTags 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromTags(Tagss []string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tags IN (?)", Tagss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ChatGroupMgr) GetFromStatus(Status int) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", Status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ChatGroupMgr) GetBatchFromStatus(Statuss []int) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", Statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ChatGroupMgr) FetchByPrimaryKey(ID string) (result ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_ChatGroupMgr) FetchByIndex(CreatorID string) (results []*ChatGroup, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("creator_id = ?", CreatorID).Find(&results).Error

	return
}
