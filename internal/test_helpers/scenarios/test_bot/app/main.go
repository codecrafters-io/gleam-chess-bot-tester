package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	// Create server with no recovery
	server := &http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/move" {
				return
			}

			fmt.Println("Received request to /move")
			request := map[string]any{}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				fmt.Printf("Error decoding request: %v\n", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Printf("Decoded request: %+v\n", request["fen"])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("e2e4"))
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
