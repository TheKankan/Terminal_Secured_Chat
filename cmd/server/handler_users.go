package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Succesfully registered")
}
