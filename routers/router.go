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
	"m4-im/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	// Setup validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.SetupValidator(v)
	}

	// Use validation error handler -> handle binding errors
	r.Use(middleware.ValidateErrorHandler(validate.Uni))

	v1Group := r.Group("api/")
	registerV1Controllers(v1Group)

	return r
}

func InitWebSocket() {

}

func registerV1Controllers(g *gin.RouterGroup) {
	api.InitUserController(g)
}
