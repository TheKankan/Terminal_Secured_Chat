package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func IsJsonValid(r *http.Response) bool {
	body, _ := io.ReadAll(r.Body)

	if r.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != "" {
			fmt.Printf("Server Error: %s\n\n", errResp.Error) // Prints JSON error message
		}
		return false
	}
	return true
}

func sendJSON(url string, payload interface{}) (*http.Response, error) {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
