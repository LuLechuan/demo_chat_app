package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

// ChatMessage represents a chat message structure.
type ChatMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

var (
	messages     []ChatMessage
	messagesLock sync.Mutex
)

// SaveChatHandler handles the saving of chat messages.
func SaveChatHandler(w http.ResponseWriter, r *http.Request) {
	var message ChatMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID (for demonstration purposes; you may want a database for this).
	message.ID = len(messages) + 1

	messagesLock.Lock()
	defer messagesLock.Unlock()

	messages = append(messages, message)

	w.WriteHeader(http.StatusCreated)
}

// ShowChatHandler handles retrieving chat messages.
func ShowChatHandler(w http.ResponseWriter, r *http.Request) {
	messagesLock.Lock()
	defer messagesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
