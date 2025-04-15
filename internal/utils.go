package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
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
	if err != nil {
		FEN += " 0 1"
		_, err = chess.FEN(FEN)
	}
	return err == nil
}

// File utils to check for the existence of entry files in the repo

func checkForEntryMd(files []os.DirEntry) bool {
	return checkForFile(files, "ENTRY.md")
}

func checkForExampleEntryMd(files []os.DirEntry) bool {
	return checkForFile(files, "example.ENTRY.md")
}

func checkForEitherExampleMdOrEntryMd(files []os.DirEntry) bool {
	return checkForExampleEntryMd(files) || checkForEntryMd(files)
}

func listFilesInExecutableDir(stageHarness *test_case_harness.TestCaseHarness) ([]os.DirEntry, error) {
	executablePath := stageHarness.Executable.Path
	executableDir := filepath.Dir(executablePath)

	files, err := os.ReadDir(executableDir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func checkForFile(files []os.DirEntry, fileName string) bool {
	if len(files) == 0 {
		return false
	}

	for _, file := range files {
		if file.Name() == fileName {
			return true
		}
	}
	return false
}
