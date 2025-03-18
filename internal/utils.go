func moveIsValid(fenStr string, move string) bool {
	fen, err := chess.FEN(fenStr)
	if err != nil {
		return false
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
