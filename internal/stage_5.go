package internal

import (
	"fmt"

	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func test5(stageHarness *test_case_harness.TestCaseHarness) error {
	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger

	files, err := listFilesInExecutableDir(stageHarness)
	if err != nil {
		return err
	}
	if !checkForEntryMd(files) {
		return fmt.Errorf("ENTRY.md not found")
	}
	logger.Successf("ENTRY.md file found! You're all set to participate in the Gleam Chess Tournament.")

	return nil
}
