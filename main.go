package main

import (
	"fmt"
	"net/http"

	"demo_chat_app/handler"
)

func main() {
	http.HandleFunc("/save", handler.SaveChatHandler)
	http.HandleFunc("/show", handler.ShowChatHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
