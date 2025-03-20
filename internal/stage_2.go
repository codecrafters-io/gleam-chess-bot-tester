package internal

import (
	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func test2(stageHarness *test_case_harness.TestCaseHarness) error {
	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	// Opening position
	FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	test_case := test_cases.GetMoveTestCase{
		FEN:                        FEN,
		AssertGeneratedMoveIsValid: false,
	}
	if err := test_case.Run(stageHarness); err != nil {
		return err
	}

	return nil
}
