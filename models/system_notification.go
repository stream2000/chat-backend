package models

import (
	"fmt"
	"time"
)

type _SystemNotificationMgr struct {
	*BaseMgr
}

// SystemNotificationMgr open func
func SystemNotificationMgr() *_SystemNotificationMgr {
	if db == nil {
		panic(fmt.Errorf("SystemNotificationMgr need init by db"))
	}
	return &_SystemNotificationMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemNotificationMgr) GetTableName() string {
	return "system_notification"
}

// Get 获取
func (obj *_SystemNotificationMgr) Get() (result SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemNotificationMgr) Gets() (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemNotificationMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithCursor cursor获取
func (obj *_SystemNotificationMgr) WithCursor(Cursor int) Option {
	return optionFunc(func(o *options) { o.query["cursor"] = Cursor })
}

// WithTitle title获取
func (obj *_SystemNotificationMgr) WithTitle(Title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = Title })
}

// WithContent content获取
func (obj *_SystemNotificationMgr) WithContent(Content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = Content })
}

// WithTimeAt time_at获取
func (obj *_SystemNotificationMgr) WithTimeAt(TimeAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_at"] = TimeAt })
}

// WithType type获取
func (obj *_SystemNotificationMgr) WithType(Type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = Type })
}

// GetByOption 功能选项模式获取
func (obj *_SystemNotificationMgr) GetByOption(opts ...Option) (result SystemNotification, err error) {
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
func (obj *_SystemNotificationMgr) GetByOptions(opts ...Option) (results []*SystemNotification, err error) {
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
func (obj *_SystemNotificationMgr) GetFromID(ID int) (result SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromID(IDs []int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromCursor 通过cursor获取内容
func (obj *_SystemNotificationMgr) GetFromCursor(Cursor int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor = ?", Cursor).Find(&results).Error

	return
}

// GetBatchFromCursor 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromCursor(Cursors []int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor IN (?)", Cursors).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_SystemNotificationMgr) GetFromTitle(Title string) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", Title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromTitle(Titles []string) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", Titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_SystemNotificationMgr) GetFromContent(Content string) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", Content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromContent(Contents []string) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", Contents).Find(&results).Error

	return
}

// GetFromTimeAt 通过time_at获取内容
func (obj *_SystemNotificationMgr) GetFromTimeAt(TimeAt time.Time) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_at = ?", TimeAt).Find(&results).Error

	return
}

// GetBatchFromTimeAt 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromTimeAt(TimeAts []time.Time) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_at IN (?)", TimeAts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_SystemNotificationMgr) GetFromType(Type int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", Type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_SystemNotificationMgr) GetBatchFromType(Types []int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", Types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemNotificationMgr) FetchByPrimaryKey(ID int) (result SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_SystemNotificationMgr) FetchByIndex(Cursor int) (results []*SystemNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor = ?", Cursor).Find(&results).Error

	return
}
