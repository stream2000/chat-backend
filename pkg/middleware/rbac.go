/*
@Time : 2020/1/18 11:25
@Author : Minus4
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"m4-im/pkg/e"
	"m4-im/pkg/response"
	"m4-im/service"
	"net/http"
)

func RbacFilter(roleCode int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rsp := response.NewResponseBuilder(c)

		groupId := c.Param("group_id")
		if groupId == "" {
			rsp.Code = e.ErrUnKnownInternalError
			rsp.Response(http.StatusInternalServerError)
			c.Abort()
			return
		}

		ok := service.ExistsGroupById(groupId)
		if !ok {
			rsp.Code = e.ErrGroupNotFound
			rsp.Response(http.StatusBadRequest)
			c.Abort()
			return
		}

		userId := c.MustGet("uid").(string)
		err := service.CheckUserRole(userId, groupId, roleCode)
		if err != nil {
			rsp.Code = err.(*service.WarpedError).Code
			rsp.Response(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
