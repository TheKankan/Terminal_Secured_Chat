package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/TheKankan/TerminalSecuredChat/internal/auth"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow terminal clients
	},
}

var (
	clients   = make(map[*websocket.Conn]string)
	clientsMu sync.Mutex
)

func (cfg *apiConfig) handlerWebSocket(w http.ResponseWriter, r *http.Request) {
	// Extract and validate JWT token
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, "missing auth token", http.StatusUnauthorized)
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	username, err := cfg.db.GetUsernameFromID(r.Context(), userID)

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	// Register client
	clientsMu.Lock()
	clients[conn] = username
	clientsMu.Unlock()

	log.Printf("Client %s connected \n", username)
	msg := []byte("User " + username + " connected")
	broadcast(websocket.TextMessage, msg)

	defer func() {
		// Defer unregister client
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		conn.Close()
		log.Printf("Client %s disconnected \n", username)
		msg := []byte("User " + username + " disconnected")
		broadcast(websocket.TextMessage, msg)
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// Format message with timestamp and username
		now := time.Now().Format("15h04")
		formattedMsg := fmt.Sprintf("[%s] %s: %s", now, username, string(msg))
		log.Printf("Received: %s\n", formattedMsg)

		// Broadcast to all clients
		broadcast(msgType, []byte(formattedMsg))
	}
}

func broadcast(msgType int, msg []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {
		err := client.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("write error:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
