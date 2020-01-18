/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"m4-im/pkg/middleware"
	"m4-im/pkg/response"
	"m4-im/service"
	"net/http"
)

func DisbandGroup(c *gin.Context) {
	rsp := response.NewResponseBuilder(c)

	groupId := c.Param("group_id")
	err := service.DeleteGroup(groupId)
	if err != nil {
		rsp.Code = err.(*service.WarpedError).Code
		rsp.Msg = err.(*service.WarpedError).Error()
		rsp.Response(http.StatusBadRequest)
		return
	}
	rsp.Response(http.StatusOK)
	return
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
	g.Use(middleware.JWT())
	{
		g.POST("/disband", middleware.RbacFilter(service.GroupAdminUser), DisbandGroup)
		g.POST("/addManager", middleware.RbacFilter(service.GroupAdminUser), AddManager)
		g.POST("/deleteManager", middleware.RbacFilter(service.GroupAdminUser), DeleteManager)
		g.POST("/allow", middleware.RbacFilter(service.GroupManagerUser), AllowMemberJoining)
		g.POST("/setInfo", middleware.RbacFilter(service.GroupManagerUser), SetGroupManagementInformation)
	}
}
