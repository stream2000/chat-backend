/*
@Time : 2020/1/15 21:54
@Author : Minus4
*/
package service

import "m4-im/models"

func Register() (err error) {
	db := models.GetDB()
	tx := db.Begin()
	defer func() {
		if nil == err {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	return
}

func Auth() (err error) {
	db := models.GetDB()
	tx := db.Begin()
	defer func() {
		if nil == err {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	return
}
