package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cfg *config) handlerLogin() bool {
	reader := bufio.NewReader(os.Stdin)
	url := "http://" + cfg.addr + "/login"

	// Get the username and check if it is valid
	var username string
	for {
		fmt.Print("Username: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidUsername(input) {
			username = input
			break
		}

		fmt.Print("❌ Invalid username (no spaces, max 20 chars)\n\n")
	}

	// Get the password and check if it is valid
	var password string
	for {
		fmt.Print("Password: ")
		input, _ := reader.ReadString('\n')

		if isValidPassword(input) {
			password = input
			break
		}

		fmt.Print("❌ Invalid password (min 4 chars, max 50 chars)\n\n")
	}

	credentials := Credentials{
		Username: username,
		Password: password,
	}
	resp, err := sendJSON(url, credentials)
	if err != nil {
		fmt.Printf("Failed to send credentials: %s \n\n", err)
		return false
	}
	defer resp.Body.Close()

	if !IsJsonValid(resp) {
		return false
	}

	// Todo : Add a way to store user struct that is being sent back

	fmt.Printf("username : %s, password : %s\n", username, password)

	return true

}

func (cfg *config) handlerRegister() bool {
	url := "http://" + cfg.addr + "/register"

	reader := bufio.NewReader(os.Stdin)

	// Get the username and check if it is valid
	var username string
	for {
		fmt.Print("Username: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidUsername(input) {
			username = input
			break
		}

		fmt.Print("❌ Invalid username (no spaces, max 20 chars)\n\n")
	}

	// Get the password and check if it is valid
	var password string
	for {
		fmt.Print("Password: ")
		input, _ := reader.ReadString('\n')

		if isValidPassword(input) {
			password = input
			break
		}

		fmt.Print("❌ Invalid password (min 4 chars, max 50 chars)\n\n")
	}

	fmt.Printf("username : %s, password : %s\n", username, password)

	credentials := Credentials{
		Username: username,
		Password: password,
	}

	resp, err := sendJSON(url, credentials)
	if err != nil {
		fmt.Printf("Failed to send credentials: %s \n\n", err)
		return false
	}
	defer resp.Body.Close()

	if !IsJsonValid(resp) {
		return false
	}

	// Todo : Add a way to store user struct that is being sent back

	fmt.Println("✅ Registration successful!")
	return true
}

func isValidUsername(s string) bool {
	const maxInput = 20

	if len(s) == 0 || len(s) > maxInput {
		return false
	}

	// No space
	for _, r := range s {
		if r == ' ' {
			return false
		}
	}

	return true
}

func isValidPassword(s string) bool {
	const minInput = 4
	const maxInput = 100

	return len(s) >= minInput && len(s) <= maxInput
}
