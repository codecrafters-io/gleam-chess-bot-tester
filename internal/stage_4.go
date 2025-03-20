package internal

import (
	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func test4(stageHarness *test_case_harness.TestCaseHarness) error {
	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	positionCounts := 4
	for i, FEN := range random.RandomElementsFromArray(WinAtChessFENs, positionCounts) {
		if !checkFEN(FEN) {
			continue
		}
		stageHarness.Logger.Infof("%d/%d RUN Generate Chess Move for Position: %s", i+1, positionCounts, FEN)

		testCase := test_cases.GetMoveTestCase{
			FEN:                        FEN,
			AssertGeneratedMoveIsValid: true,
		}
		if err := testCase.Run(stageHarness); err != nil {
			return err
		}
	}
	return nil
}
