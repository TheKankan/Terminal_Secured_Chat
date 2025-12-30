package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func (cfg *config) handlerWebSocket() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
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
			fmt.Println("Server:", string(msg))
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
