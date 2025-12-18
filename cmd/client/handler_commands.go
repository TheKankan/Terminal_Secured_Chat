package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cfg *apiConfig) handlerLogin() {
	reader := bufio.NewReader(os.Stdin)

	// Get the username and check if it is valid
	var username string
	for {
		fmt.Print("Username: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidInput(input, 20) {
			username = input
			break
		}

		fmt.Println("❌ Invalid username (no spaces, max 20 chars)\n")
	}

	// Get the password and check if it is valid
	var password string
	for {
		fmt.Print("Password: ")
		bytePwd, _ := term.ReadPassword(int(syscall.Stdin))
		input := strings.TrimSpace(string(bytePwd))

		if isValidInput(input, 50) {
			password = input
			break
		}

		fmt.Println("❌ Invalid password (no spaces, max 50 chars)\n")
	}

	credentials := Credentials{
		Username: username,
		Password: password,
	}
	_, err := sendJSON(cfg.addr, credentials)
	if err != nil {
		fmt.Println("Failed to send credentials:", err)
		return
	}

}

func (cfg *apiConfig) handlerRegister() {

}

func isValidInput(s string, maxInput int) bool {
	if len(s) == 0 || len(s) > maxInput {
		return false
	}

	// No space or multiple lines
	for _, r := range s {
		if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
			return false
		}
	}

	return true
}
