package model

import (
	"time"
)

// ChatGroup [...]
type ChatGroup struct {
	ID        string    `gorm:"primary_key;column:id;type:varchar(256);not null" json:"-"`
	GroupName string    `gorm:"column:group_name;type:varchar(256);not null" json:"group_name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	CreatorID string    `gorm:"index;column:creator_id;type:varchar(256);not null" json:"creator_id"`
	Info      string    `gorm:"column:info;type:varchar(256)" json:"info"`
	Tags      string    `gorm:"column:tags;type:varchar(1024)" json:"tags"`
	Status    int       `gorm:"column:status;type:int(11)" json:"status"`
}

// GroupTag [...]
type GroupTag struct {
	GroupID string `gorm:"primary_key;column:group_id;type:varchar(256);not null" json:"group_id"`
	TagID   int    `gorm:"primary_key;column:tag_id;type:int(11);not null" json:"tag_id"`
}

// GroupAdmin [...]
type GroupAdmin struct {
	UserID   string `gorm:"primary_key;column:user_id;type:varchar(256);not null" json:"user_id"`
	GroupID  string `gorm:"primary_key;column:group_id;type:varchar(256);not null" json:"group_id"`
	RoleCode int    `gorm:"column:role_code;type:int(3)" json:"role_code"`
}

// Tag [...]
type Tag struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name      string    `gorm:"unique;column:name;type:varchar(256)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
}

// UserNotification [...]
type UserNotification struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cursor       int       `gorm:"index:user_notification_id_cursor_index;column:cursor;type:int(19)" json:"cursor"`
	Title        string    `gorm:"column:title;type:varchar(256);not null" json:"title"`
	Content      string    `gorm:"column:content;type:varchar(256)" json:"content"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	Type         int       `gorm:"column:type;type:int(11)" json:"type"`
	PublisherID  string    `gorm:"index;column:publisher_id;type:varchar(256);not null" json:"publisher_id"`
	SubscriberID string    `gorm:"index;column:subscriber_id;type:varchar(256);not null" json:"subscriber_id"`
}

// Membership [...]
type Membership struct {
	GroupID   string    `gorm:"primary_key;column:group_id;type:varchar(256);not null" json:"group_id"`
	UserID    string    `gorm:"primary_key;column:user_id;type:varchar(256);not null" json:"user_id"`
	AdminID   string    `gorm:"index;column:admin_id;type:varchar(256);not null" json:"admin_id"` // 允许进群的管理员的id
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
}

// SysRole [...]
type SysRole struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name      string `gorm:"column:name;type:varchar(30)" json:"name"`
	CreatorID string `gorm:"index;column:creator_id;type:varchar(256)" json:"creator_id"`
}

// SystemAdmin [...]
type SystemAdmin struct {
	UserID string `gorm:"primary_key;column:user_id;type:varchar(256);not null" json:"user_id"`
	RoleID int    `gorm:"primary_key;column:role_id;type:int(11);not null" json:"role_id"`
}

// SystemNotification [...]
type SystemNotification struct {
	ID      int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cursor  int       `gorm:"index;column:cursor;type:int(19)" json:"cursor"`
	Title   string    `gorm:"column:title;type:varchar(256);not null" json:"title"`
	Content string    `gorm:"column:content;type:varchar(256)" json:"content"`
	TimeAt  time.Time `gorm:"column:time_at;type:datetime" json:"time_at"`
	Type    int       `gorm:"column:type;type:int(3);not null" json:"type"`
}

// User [...]
type User struct {
	ID        string    `gorm:"primary_key;column:id;type:varchar(256);not null" json:"-"`
	Email     string    `gorm:"unique;column:email;type:varchar(256)" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(256);not null" json:"password"`
	Nickname  string    `gorm:"column:nickname;type:varchar(256)" json:"nickname"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	Status    int       `gorm:"column:status;type:int(11)" json:"status"`
}
