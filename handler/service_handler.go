package handler

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ChatMessage represents a chat message structure.
type ChatMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type EmptyResp struct{}

var (
	messages     []ChatMessage
	messagesLock sync.Mutex
)

func SaveChatHertzHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var message ChatMessage
	err = c.Bind(&message)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	// Assign a unique ID (for demonstration purposes; you may want a database for this).
	message.ID = len(messages) + 1

	messagesLock.Lock()
	defer messagesLock.Unlock()

	messages = append(messages, message)

	c.Status(consts.StatusOK)
	body, _ := json.Marshal(&EmptyResp{})
	c.Write(body)
}

func ShowChatHertzHandler(ctx context.Context, c *app.RequestContext) {
	messagesLock.Lock()
	defer messagesLock.Unlock()
	c.Header("Content-Type", "application/json")
	body, err := json.Marshal(messages)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.Write(body)
}
