/*
@Time : 2020/1/15 21:02
@Author : Minus4
*/
package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"m4-im/pkg/setting"
)

var (
	ErrRecordNotFound = errors.New("记录不存在")
	ErrRecordDeleted  = errors.New("记录已被删除")
	ErrRecordExists   = errors.New("记录已存在")
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("dao.Setup err: %v", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func ReConnect() {
	Setup()
}

func GetDB() *gorm.DB {
	return db
}
func RawDB() *sql.DB {
	return db.DB()
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
