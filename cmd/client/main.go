package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := "8080"
	addr := "localhost:" + port

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected")

}
