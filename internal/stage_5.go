package internal

import (
	"fmt"
	"os"
	"path/filepath"

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

func checkForEntryMd(files []os.DirEntry) bool {
	return checkForFile(files, "ENTRY.md")
}

func checkForExampleEntryMd(files []os.DirEntry) bool {
	return checkForFile(files, "example.ENTRY.md")
}

func checkForEitherExampleOrEntryMd(files []os.DirEntry) bool {
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
