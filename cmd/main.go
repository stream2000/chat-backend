/*
@Time : 2020/1/15 23:32
@Author : Minus4
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"m4-im/controllers"
	"m4-im/controllers/websocket"
	"m4-im/dao"
	"m4-im/pkg/middleware"
	"m4-im/pkg/setting"
	"m4-im/pkg/util"
	"net/http"
)

func init() {
	util.Setup()
	dao.Setup()
	websocket.SetUpChannelManager()
}

func ServeHttp() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := controllers.InitHttpRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s\n", endPoint)
	server.ListenAndServe()
}

func ServeWebsocket() {
	router := gin.New()
	server := controllers.InitWebSocketRouter()
	go server.Serve()
	defer server.Close()

	router.Use(middleware.GinMiddleware())
	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/", gin.WrapH(server))
	router.Run(fmt.Sprintf(":%d", setting.ServerSetting.WebSocketPort))
}
func main() {
	go ServeHttp()
	ServeWebsocket()
}
