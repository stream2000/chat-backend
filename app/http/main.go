/*
@Time : 2020/1/15 23:32
@Author : Minus4
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"m4-im/pkg/setting"
	rbac "m4-im/rbac/util"
	"m4-im/routers"
	"net/http"
)

func init() {
	setting.Setup()
	rbac.Setup()
}
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
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

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}