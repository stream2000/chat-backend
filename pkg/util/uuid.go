/*
@Time : 2020/1/16 17:15
@Author : Minus4
*/
package util

import (
	"github.com/satori/go.uuid"
)

func UUID() string {
	return uuid.NewV4().String()
}
