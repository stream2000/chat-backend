package models

import (
	"fmt"
	"time"
)

type _UserNotificationMgr struct {
	*BaseMgr
}

// UserNotificationMgr open func
func UserNotificationMgr() *_UserNotificationMgr {
	if db == nil {
		panic(fmt.Errorf("UserNotificationMgr need init by db"))
	}
	return &_UserNotificationMgr{BaseMgr: &BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserNotificationMgr) GetTableName() string {
	return "user_notification"
}

// Get 获取
func (obj *_UserNotificationMgr) Get() (result UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserNotificationMgr) Gets() (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserNotificationMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithCursor cursor获取
func (obj *_UserNotificationMgr) WithCursor(Cursor int) Option {
	return optionFunc(func(o *options) { o.query["cursor"] = Cursor })
}

// WithTitle title获取
func (obj *_UserNotificationMgr) WithTitle(Title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = Title })
}

// WithContent content获取
func (obj *_UserNotificationMgr) WithContent(Content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = Content })
}

// WithCreatedAt created_at获取
func (obj *_UserNotificationMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// WithType type获取
func (obj *_UserNotificationMgr) WithType(Type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = Type })
}

// WithPublisherID publisher_id获取
func (obj *_UserNotificationMgr) WithPublisherID(PublisherID string) Option {
	return optionFunc(func(o *options) { o.query["publisher_id"] = PublisherID })
}

// WithSubscriberID subscriber_id获取
func (obj *_UserNotificationMgr) WithSubscriberID(SubscriberID string) Option {
	return optionFunc(func(o *options) { o.query["subscriber_id"] = SubscriberID })
}

// GetByOption 功能选项模式获取
func (obj *_UserNotificationMgr) GetByOption(opts ...Option) (result UserNotification, err error) {
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
func (obj *_UserNotificationMgr) GetByOptions(opts ...Option) (results []*UserNotification, err error) {
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
func (obj *_UserNotificationMgr) GetFromID(ID int) (result UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromID(IDs []int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromCursor 通过cursor获取内容
func (obj *_UserNotificationMgr) GetFromCursor(Cursor int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor = ?", Cursor).Find(&results).Error

	return
}

// GetBatchFromCursor 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromCursor(Cursors []int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor IN (?)", Cursors).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_UserNotificationMgr) GetFromTitle(Title string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", Title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromTitle(Titles []string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", Titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_UserNotificationMgr) GetFromContent(Content string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", Content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromContent(Contents []string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", Contents).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_UserNotificationMgr) GetFromCreatedAt(CreatedAt time.Time) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_UserNotificationMgr) GetFromType(Type int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", Type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromType(Types []int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", Types).Find(&results).Error

	return
}

// GetFromPublisherID 通过publisher_id获取内容
func (obj *_UserNotificationMgr) GetFromPublisherID(PublisherID string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("publisher_id = ?", PublisherID).Find(&results).Error

	return
}

// GetBatchFromPublisherID 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromPublisherID(PublisherIDs []string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("publisher_id IN (?)", PublisherIDs).Find(&results).Error

	return
}

// GetFromSubscriberID 通过subscriber_id获取内容
func (obj *_UserNotificationMgr) GetFromSubscriberID(SubscriberID string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscriber_id = ?", SubscriberID).Find(&results).Error

	return
}

// GetBatchFromSubscriberID 批量唯一主键查找
func (obj *_UserNotificationMgr) GetBatchFromSubscriberID(SubscriberIDs []string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscriber_id IN (?)", SubscriberIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserNotificationMgr) FetchByPrimaryKey(ID int) (result UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUserNotificationIDCursorIndexIndex  获取多个内容
func (obj *_UserNotificationMgr) FetchByUserNotificationIDCursorIndexIndex(Cursor int) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cursor = ?", Cursor).Find(&results).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_UserNotificationMgr) FetchByIndex(PublisherID string, SubscriberID string) (results []*UserNotification, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("publisher_id = ? AND subscriber_id = ?", PublisherID, SubscriberID).Find(&results).Error

	return
}
