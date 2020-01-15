/*
@Time : 2020/1/16 00:58
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
)

func SendSystemNotification(c *gin.Context) {

}

func PullSystemNotification(c *gin.Context) {

}

func SendUserNotification(c *gin.Context) {

}

func PullUserNotification(c *gin.Context) {

}

func InitNotificationController(g *gin.RouterGroup) {
	g.Use()
	{
		g.GET("/sys", PullSystemNotification)
		g.GET("/user", PullUserNotification)
		g.POST("/sys", SendSystemNotification)
		g.POST("/user", SendUserNotification)
	}
}
