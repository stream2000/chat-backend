/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"m4-im/middleware"
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

	// Init group Admin Controller
	groupAdminController := r.Group("api/v1/group/admin/:group_id")
	apiV1.InitGroupAdminController(groupAdminController)

	// Init group Controller
	groupController := r.Group("api/v1/group")
	apiV1.InitGroupController(groupController)

	// Init user Controller
	userController := r.Group("api/v1/user")
	apiV1.InitUserController(userController)

	// Init notification Controller
	notificationController := r.Group("api/v1/notify")
	apiV1.InitNotificationController(notificationController)

	return r
}
