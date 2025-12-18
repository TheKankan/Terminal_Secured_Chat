package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TheKankan/TerminalSecuredChat/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	db        *database.Queries
	jwtSecret string
	addr      string
}

func main() {
	const port = "8080"

	// Getting .env variables
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("SECRET environment variable is not set")
	}

	// Setting up connection to postgres database
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	// Setting adress
	addr := "localhost:" + port

	// Saving variables in config
	apiCfg := apiConfig{
		db:        dbQueries,
		jwtSecret: jwtSecret,
		addr:      addr,
	}

	reader := bufio.NewReader(os.Stdin)

	// Login or Registering initial state
	for {
		fmt.Print("Login / Register : ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "login" {
			fmt.Print("User chose LOGIN\n\n")
			// appeler la logique login
			break
		}
		if input == "register" {
			fmt.Print("User chose REGISTER\n\n")
			// appeler la logique register
			break
		}
		fmt.Print("Invalid command. Please type 'login' or 'register'\n\n")
	}

	// User can send & receive messages
	fmt.Print("Welcome [User] ! You are now connected to the chat ! \n\n")
	input, _ := reader.ReadString('\n')
	fmt.Printf("%s", input)

}
