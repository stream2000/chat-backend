/*
@Time : 2020/1/17 00:48
@Author : Minus4
*/
package util

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

//email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func AbortWithBindError(c *gin.Context, err error) {
	_ = c.Error(err).SetType(gin.ErrorTypeBind)
	c.Abort()
}

func AbortWithInvalidBasicAuthError(c *gin.Context) {

}
