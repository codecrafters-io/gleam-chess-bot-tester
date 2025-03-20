package assertions

import (
	"fmt"
	"io"
	"net/http"

	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/corentings/chess"
)

type ValidMoveAssertion struct {
	FEN string
}

func (a *ValidMoveAssertion) Run(response *http.Response, logger *logger.Logger) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %v", err)
	}

	move := string(body)
	if !moveIsValid(a.FEN, move) {
		return fmt.Errorf("Invalid move received: %s", move)
	}
	logger.Successf("✓ Received valid move %s", move)

	return nil
}

func moveIsValid(fenStr string, move string) bool {
	var fen func(*chess.Game)
	var err error

	fen, err = chess.FEN(fenStr)
	if err != nil {
		fenStr += " 0 1"
		fen, err = chess.FEN(fenStr)
		if err != nil {
			return false
		}
	}

	game := chess.NewGame(fen)

	// Try to make the move
	err = game.MoveStr(move)
	if err != nil {
		// If move is invalid, try UCI notation
		moveUCI, err := chess.UCINotation{}.Decode(game.Position(), move)
		if err != nil {
			return false
		}

		// Check if the UCI move is in the list of valid moves
		validMoves := game.ValidMoves()
		for _, validMove := range validMoves {
			if moveUCI.String() == validMove.String() {
				return true
			}
		}
		return false
	}

	return true
}
