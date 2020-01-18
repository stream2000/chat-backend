/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/pkg/e"
	"m4-im/pkg/middleware"
	"m4-im/pkg/response"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
	"m4-im/service"
	"net/http"
)

func Register(c *gin.Context) {
	rsp := response.NewResponseBuilder(c)

	var param params.RegisterParam
	if err := c.ShouldBindJSON(&param); err == nil {
		err = service.Register(param)
		if err != nil {
			rsp.Code = err.(*service.WarpedError).Code
			rsp.Msg = err.(*service.WarpedError).Error()
			rsp.Response(http.StatusBadRequest)
			return
		}
		rsp.Response(http.StatusOK)
	} else {
		util.AbortWithBindError(c, err)
		return
	}
}

func Auth(c *gin.Context) {
	rsp := response.NewResponseBuilder(c)

	basicAuthHeader := c.GetHeader("Authorization")
	email, password, err := util.ParseBasicHeader(basicAuthHeader)
	if err != nil {
		rsp.Code = e.ErrInvalidBasicAuthParam
		rsp.Msg = err.Error()
		rsp.Response(http.StatusUnauthorized)
		return
	}

	ok, id := service.Auth(email, password)
	if !ok {
		rsp.Code = e.ErrBasicAuthFailed
		rsp.Response(http.StatusUnauthorized)
		return
	}

	token, err := util.GenerateToken(email, password, id)
	if err != nil {
		rsp.Code = e.ErrJwtAuth
		rsp.Response(http.StatusInternalServerError)
		return
	}
	rsp.Data = map[string]interface{}{
		"token": token,
	}
	rsp.Response(http.StatusOK)
}

func SetBasicInfo(c *gin.Context) {
	rsp := response.NewResponseBuilder(c)
	rsp.Code = e.SUCCESS
	idValue, _ := c.Get("uid")
	rsp.Data = idValue
	rsp.Response(http.StatusOK)
}

func InitUserController(g *gin.RouterGroup) {
	g.GET("/auth", Auth)
	g.POST("/register", Register)
	g.POST("/setInfo", middleware.JWT(), SetBasicInfo)
}
