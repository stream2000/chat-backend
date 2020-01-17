/*
@Time : 2020/1/17 00:11
@Author : Minus4
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"m4-im/pkg/e"
	"net/http"
)

type validateErrorHandler struct {
	uni *ut.UniversalTranslator
}

func ValidateErrorHandler(uni *ut.UniversalTranslator) gin.HandlerFunc {
	return newErrorHandler(uni).HandleErrors
}

func newErrorHandler(uni *ut.UniversalTranslator) *validateErrorHandler {
	return &validateErrorHandler{
		uni: uni,
	}
}

func (h *validateErrorHandler) HandleErrors(c *gin.Context) {
	// a post middleware
	c.Next()
	//if len(c.Errors) > 0 {
	//	for _, e := range c.Errors {
	//	// Find out what type of error it is
	//	switch e.Type {
	//	case gin.ErrorTypePublic:
	//		// Only output public errors if nothing has been written yet
	//		if !c.Writer.Written() {
	//			c.JSON(c.Writer.Status(), gin.H{"Error": e.Error()})
	//		}
	//	case gin.ErrorTypeBind:
	//
	//
	//	default:
	//	}
	//}
	//}
	errorToPrint := c.Errors.ByType(gin.ErrorTypeBind).Last()
	if errorToPrint != nil {
		if errs, ok := errorToPrint.Err.(validator.ValidationErrors); ok {
			trans, _ := h.uni.GetTranslator("en") // 这里也可以通过获取 HTTP Header 中的 Accept-Language 来获取用户的语言设置
			c.JSON(http.StatusBadRequest, gin.H{
				"code": e.ErrInvalidParams,
				"msg":  errs.Translate(trans),
			})
			return
		}
		// deal with other errors ...
	}
}
