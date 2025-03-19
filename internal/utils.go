package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/corentings/chess"
)

func makeMove(fenStr string) (string, error) {
	fen, err := chess.FEN(fenStr)
	if err != nil {
		return "", err
	}

	game := chess.NewGame(fen)

	// Determine turn color
	turn := "white"
	if game.Position().Turn() == chess.Black {
		turn = "black"
	}

	// Prepare request body
	type requestBody struct {
		Fen         string   `json:"fen"`
		Turn        string   `json:"turn"`
		FailedMoves []string `json:"failed_moves"`
	}

	body := requestBody{
		Fen:         fenStr,
		Turn:        turn,
		FailedMoves: []string{},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("error marshaling request body: %v", err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(
		"http://localhost:8000/move",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to make move, got status %d: %s", resp.StatusCode, string(responseBody))
	}

	return string(responseBody), nil
}

func checkFEN(FEN string) bool {
	_, err := chess.FEN(FEN)
	return err == nil
}
