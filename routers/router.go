/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"m4-im/pkg/middleware"
	"m4-im/pkg/validate"
	apiV1 "m4-im/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Setup validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.SetupValidator(v)
	}

	// Use validation error handler -> handle binding errors
	r.Use(middleware.ValidateErrorHandler(validate.Uni))

	v1Group := r.Group("api/v1")
	registerV1Controllers(v1Group)

	return r
}

func registerV1Controllers(g *gin.RouterGroup) {
	// Init group Admin Controller
	groupAdminController := g.Group("group/admin/:group_id")
	apiV1.InitGroupAdminController(groupAdminController)

	// Init group Controller
	groupController := g.Group("group")
	apiV1.InitGroupController(groupController)

	// Init user Controller
	userController := g.Group("user")
	apiV1.InitUserController(userController)

	// Init notification Controller
	notificationController := g.Group("notify")
	apiV1.InitNotificationController(notificationController)
}
