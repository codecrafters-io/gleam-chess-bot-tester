package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server...")
	os.Stdout.Sync()
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request to /move")
		os.Stdout.Sync()
		request := map[string]any{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			fmt.Printf("Error decoding request: %v\n", err)
			os.Stdout.Sync()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Decoded request: %+v\n", request)
		os.Stdout.Sync()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("e2e4"))
	})

	http.ListenAndServe(":8000", nil)
}
