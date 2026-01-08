package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func (cfg *config) handlerWebSocket() {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+cfg.token)

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", header)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	// receive messages
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				return
			}
			// Print received message
			fmt.Println(string(msg))
		}
	}()

	// send terminal input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type messages and press enter:")

	for scanner.Scan() {
		text := scanner.Text()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("write error:", err)
			return
		}
	}
}
