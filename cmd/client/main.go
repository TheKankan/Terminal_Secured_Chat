package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type config struct {
	addr string
}

func main() {
	port := "8080"
	host := "localhost"

	addr := host + ":" + port

	// Saving variables in config
	cfg := config{
		addr: addr,
	}

	// Read first input
	reader := bufio.NewReader(os.Stdin)

	// Login or Registering initial state
	for {
		fmt.Print("Login / Register : ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "login" {
			fmt.Print("User chose LOGIN\n\n")
			if cfg.handlerLogin() {
				break
			}
		} else if input == "register" {
			fmt.Print("User chose REGISTER\n\n")
			if cfg.handlerRegister() {
				break
			}
		} else {
			fmt.Print("Invalid command. Please type 'login' or 'register'\n\n")
		}
	}

	// User can send & receive messages
	fmt.Print("Welcome [User] ! You are now connected to the chat ! \n\n")
	input, _ := reader.ReadString('\n')
	fmt.Printf("%s", input)

}
