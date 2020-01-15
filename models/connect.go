/*
@Time : 2020/1/15 21:02
@Author : Minus4
*/
package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"m4-im/pkg/setting"
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
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func ReConnect() {
	Setup()
}

func GetDB() *gorm.DB {
	return db
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
