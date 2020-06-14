/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package web

import (
	"github.com/gin-gonic/gin"
	"m4-im/controllers/websocket"
	"m4-im/dao"
	"m4-im/pkg/e"
	"m4-im/pkg/middleware"
	"m4-im/pkg/response"
	"m4-im/pkg/setting"
	"m4-im/pkg/util"
	"net/http"
	"strconv"
)

func Auth(c *gin.Context) {
	rsp := response.NewResponseBuilder(c)
	var param AuthParam
	if err := c.ShouldBindJSON(&param); err != nil {
		util.AbortWithBindError(c, err)
		return
	}
	db := dao.GetDB()

	u := &dao.User{
		Account: param.Account,
	}
	db.Table("user").Where("account = ?", param.Account).First(&u)
	if u.Id <= 0 {
		if !setting.ServerSetting.EnableRegistration {
			rsp.SetCode(e.ErrBasicAuthFailed).OK()
			return
		}
		u.PassWd = param.Password
		u.AvatarUrl = "/static/3.jpg"
		db.Create(u)
		websocket.AddNewUser(*u)
	} else {
		if u.PassWd != param.Password {
			rsp.SetCode(e.ErrBasicAuthFailed).OK()
			return
		}
	}
	token, err := util.GenerateToken(u.Account, u.PassWd, u.Id)
	if err != nil {
		rsp.Code = e.ErrJwtAuth
		rsp.Response(http.StatusInternalServerError)
		return
	}
	rsp.Data = map[string]interface{}{
		"jwt":        token,
		"avatar_url": u.AvatarUrl,
		"id":         u.Id,
	}
	rsp.Response(http.StatusOK)
}

func GetBasicInfo(c *gin.Context) {
	db := dao.GetDB()

	rsp := response.NewResponseBuilder(c)
	idValue, _ := c.Get("uid")
	userId, _ := strconv.ParseInt(idValue.(string), 10, 64)
	user := &dao.User{}

	db.First(user, int(userId))
	otherUsers := make([]*dao.User, 0)
	db.Table("user").Not("id = ?", userId).Find(&otherUsers)

	for _, u := range otherUsers {
		u.PassWd = ""
	}

	data := map[string]interface{}{
		"user": map[string]interface{}{
			"name": user.Account,
			"img":  user.AvatarUrl,
		},
		"sessions": otherUsers,
	}
	rsp.Data = data
	rsp.Response(http.StatusOK)
}

func InitUserController(g *gin.RouterGroup) {
	g.POST("/auth", Auth)
	g.GET("/init", middleware.JWT(), GetBasicInfo)
}
