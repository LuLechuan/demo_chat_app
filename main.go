package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"

	"demo_chat_app/handler"
)

func main() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.GET("/show", handler.ShowChatHertzHandler)
	h.POST("/save", handler.SaveChatHertzHandler)
	//http.HandleFunc("/save", handler.SaveChatHandler)
	//http.HandleFunc("/show", handler.ShowChatHandler)

	fmt.Println("Server is running on :8080")
	h.Spin()
}
