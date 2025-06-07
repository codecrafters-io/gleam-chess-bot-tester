package internal

import (
	"fmt"

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

	logger := stageHarness.Logger

	positionCounts := 3
	for i, FEN := range random.RandomElementsFromArray(WinAtChessFENs, positionCounts) {
		if !checkFEN(FEN) {
			continue
		}

		testCase := test_cases.GetMoveTestCase{
			FEN:                        FEN,
			AssertGeneratedMoveIsValid: true,
		}
		logger.UpdateSecondaryPrefix(fmt.Sprintf("board-%d", i+1))
		stageHarness.Logger.Infof("$ %s", getCurlCommandForRequest(FEN))

		if err := testCase.Run(stageHarness, logger); err != nil {
			return err
		}
		logger.ResetSecondaryPrefix()
		logger.Successf("Successfully generated move for position %d", i+1)
	}
	return nil
}
