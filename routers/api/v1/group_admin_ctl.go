/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/rbac/middleware/jwt"
)

func DisbandGroup(c *gin.Context) {

}

func AddManager(c *gin.Context) {

}

func DeleteManager(c *gin.Context) {

}

func AllowMemberJoining(c *gin.Context) {

}

// uncertain function
func SetGroupManagementInformation(c *gin.Context) {

}

func InitGroupAdminController(g *gin.RouterGroup) {
	g.Use(jwt.JWT())
	{
		g.POST("/disband", DisbandGroup)
		g.POST("/addManager", AddManager)
		g.POST("/deleteManager", DeleteManager)
		g.POST("/allow", AllowMemberJoining)
		g.POST("/setInfo", SetGroupManagementInformation)
	}
}
