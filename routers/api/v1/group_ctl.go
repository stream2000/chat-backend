/*
@Time : 2020/1/15 22:29
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/dao"
	"m4-im/pkg/middleware"
	"m4-im/pkg/response"
	"m4-im/pkg/util"
	"m4-im/pojo/params"
	"m4-im/service"
	"net/http"
)

func CreateGroup(c *gin.Context) {
	tx := dao.GetDB().Begin()
	var err error
	defer func() {
		if nil == err {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	rsp := response.NewResponseBuilder(c)
	userId := c.MustGet("uid").(string)
	var param params.AddNewGroupParam
	if err := c.ShouldBindJSON(&param); err == nil {
		err = service.CreatGroup(param, userId)
		if err != nil {
			rsp.Code = err.(*service.WarpedError).Code
			rsp.Msg = err.(*service.WarpedError).Error()
			rsp.Response(http.StatusBadRequest)
			return
		}
	} else {
		util.AbortWithBindError(c, err)
		return
	}
	rsp.Response(http.StatusOK)
}

func GetGroupInfo(c *gin.Context) {

}

func ParticipateGroup(c *gin.Context) {

}

func GetAllGroups(c *gin.Context) {

}

func GetAllGroupsByMultiTags(c *gin.Context) {

}

func InitGroupController(g *gin.RouterGroup) {
	g.Use(middleware.JWT())
	{
		g.GET("/info", GetGroupInfo)
		g.GET("/all", GetAllGroups)
		g.POST("/all/tags", GetAllGroupsByMultiTags)
		g.POST("/create", CreateGroup)
		g.POST("/participate", ParticipateGroup)
	}
}
