/*
@Time : 2020/1/15 21:53
@Author : Minus4
*/
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	"github.com/googollee/go-socket.io"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"m4-im/controllers/web"
	"m4-im/pkg/middleware"
	"m4-im/pkg/validate"
	"net/http"
)

func InitHttpRouter() *gin.Engine {
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
	web.InitUserController(v1Group)

	return r
}

func InitWebSocketRouter() *socketio.Server {
	pt := polling.Default
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return server
}
