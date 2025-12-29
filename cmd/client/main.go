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

	var user *User

	// Login or Registering initial state
	for {
		fmt.Print("Login / Register : ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "login" {
			fmt.Print("User chose LOGIN\n\n")
			// If login successful, break loop
			if u := cfg.handlerLogin(); u != nil {
				user = u
				break
			}
		} else if input == "register" {
			fmt.Print("User chose REGISTER\n\n")
			// If register successful, break loop
			if u := cfg.handlerRegister(); u != nil {
				user = u
				break
			}
		} else {
			fmt.Print("Invalid command. Please type 'login' or 'register'\n\n")
		}
	}

	// User can send & receive messages
	fmt.Printf("Welcome %s ! You are now connected to the chat ! \n\n", user.Username)
	for {
		input, _ := reader.ReadString('\n')
		fmt.Printf("%s", input)

		// TODO : send messages when the user writes something

		// TODO : listen for the server in case it sends messages back
	}

}
