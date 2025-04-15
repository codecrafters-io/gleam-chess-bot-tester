package internal

import (
	"os"
	"path/filepath"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
	"github.com/corentings/chess"
)

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
