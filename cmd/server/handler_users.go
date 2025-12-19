package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Succesfully registered")

	// get the context to pass to the query = r.Context()
	// A compléter
}
