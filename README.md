# Terminal_Secured_Chat

This project allows multiple users to communicate via a chat in a terminal

## Motivation

I love Discord but you never know if the message you just sent is 100% private.

When i want to send the netflix password to my girlfriend or share private informations I want to have my own way to do it without my informations potentially getting sold, that's why i made this project.

This project also gave me the opportunity to practice building a functional backend and managing an HTML server from scratch.

## Quick Start

### Prerequisites
- Go Version 1.25+ installed, you can download it [here](https://go.dev/doc/install)

### Installing

## Usage

## Contributing

- Clone the repo 
- Replace the variables in .env.example with your values and rename it .env (you will need to create your postgres database)
- Migrate up everything in sql/schema
- Use `go run ./cmd/server` to run the server and then `go run ./cmd/client` in another terminal for each client
