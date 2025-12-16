package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	port := "8080"
	addr := "localhost:" + port

	fmt.Println("Hello, Server!")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())
		conn.Close() // temporary
	}
}
