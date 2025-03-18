package main

import "net/http"

func main() {
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("expected response here"))
	})

	http.ListenAndServe(":8000", nil)
}
