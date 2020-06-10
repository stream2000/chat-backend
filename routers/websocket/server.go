/*
@Time : 2020/6/10 10:58
@Author : Minus4
*/
package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"log"
	"m4-im/pkg/setting"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

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

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice ", msg)
		param := &struct {
			Name string `json:"name"`
			Jwt  string `json:"jwt"`
		}{}
		err := json.Unmarshal([]byte(msg), param)
		if err != nil {
			fmt.Println(err)
		} else {
			s.Emit("reply", param)
		}
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		fmt.Println("here ", msg)
		param := &struct {
			Name string `json:"name"`
			Jwt  string `json:"jwt"`
		}{}
		err := json.Unmarshal([]byte(msg), param)
		if err != nil {
			fmt.Println(err)
		} else {
			s.Emit("reply", param)
		}
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	mux.Handle("/socket.io/", server)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})

	// decorate existing handler with cors functionality set in c
	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", setting.ServerSetting.WebSocketPort), handler))
}
