/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/rbac/middleware/jwt"
)

func Register(c *gin.Context) {

}

func Auth(c *gin.Context) {

}

func SetBasicInfo(c *gin.Context) {

}

func InitUserController(g *gin.RouterGroup) {
	g.GET("/auth", Auth)
	g.POST("/register", Register)
	g.POST("/setInfo", jwt.JWT(), SetBasicInfo)
}
