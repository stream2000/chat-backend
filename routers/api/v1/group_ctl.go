/*
@Time : 2020/1/15 22:29
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/middleware"
)

func CreateGroup(c *gin.Context) {

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
