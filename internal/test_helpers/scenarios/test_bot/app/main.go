package main

import (
	"encoding/json"
	"net/http"

	"github.com/corentings/chess"
)

func main() {
	server := &http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/move" {
				return
			}

			// fmt.Println("Received request to /move")
			request := map[string]any{}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fen, err := chess.FEN(request["fen"].(string))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			game := chess.NewGame(fen)
			moves := game.ValidMoves()

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(moves[0].String()))
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
