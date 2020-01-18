package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMgr) WithID(ID string) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithEmail email获取
func (obj *_UserMgr) WithEmail(Email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = Email })
}

// WithPassword password获取
func (obj *_UserMgr) WithPassword(Password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = Password })
}

// WithNickname nickname获取
func (obj *_UserMgr) WithNickname(Nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = Nickname })
}

// WithCreatedAt created_at获取
func (obj *_UserMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// WithStatus status获取
func (obj *_UserMgr) WithStatus(Status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = Status })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
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
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
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
func (obj *_UserMgr) GetFromID(ID string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromID(IDs []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_UserMgr) GetFromEmail(Email string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", Email).Find(&result).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromEmail(Emails []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email IN (?)", Emails).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UserMgr) GetFromPassword(Password string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password = ?", Password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromPassword(Passwords []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password IN (?)", Passwords).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容
func (obj *_UserMgr) GetFromNickname(Nickname string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname = ?", Nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromNickname(Nicknames []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname IN (?)", Nicknames).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_UserMgr) GetFromCreatedAt(CreatedAt time.Time) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_UserMgr) GetFromStatus(Status int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", Status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromStatus(Statuss []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", Statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(ID string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_UserMgr) FetchByUnique(Email string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", Email).Find(&result).Error

	return
}
