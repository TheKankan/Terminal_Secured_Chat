package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	const port = "8080"

	// Loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	addr := "localhost:" + port

	// A modifier pour HTTP
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Demander Login ou register

	// Login
	// Register

	// Une fois que le user est logged in : Lui permettre d'envoyer des messages
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println(username)

}
