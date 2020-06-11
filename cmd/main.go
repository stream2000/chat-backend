/*
@Time : 2020/1/15 23:32
@Author : Minus4
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"log"
	"m4-im/controllers"
	"m4-im/dao"
	"m4-im/pkg/setting"
	"m4-im/pkg/util"
	"net/http"
)

func init() {
	util.Setup()
	dao.Setup()
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
	mux := http.NewServeMux()
	server := controllers.InitWebSocketRouter()
	go server.Serve()
	defer server.Close()
	mux.Handle("/socket.io/", server)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)
	log.Printf("[info] start websockrt server listening %d\n", setting.ServerSetting.WebSocketPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", setting.ServerSetting.WebSocketPort), handler))
}
func main() {
	go ServeHttp()
	ServeWebsocket()
}
