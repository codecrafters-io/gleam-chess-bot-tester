package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	loggerUtils "github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func antiCheatExecute(stageHarness *test_case_harness.TestCaseHarness) error {
	logger := stageHarness.Logger

	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	// If we can't run the executable, it must be an internal error.
	if err := b.Run(); err != nil {
		logger.Criticalf("CodeCrafters internal error. Error instantiating executable: %v", err)
		logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
		return fmt.Errorf("anti-cheat (ac1) failed")
	}

	return nil
}

func antiCheatDeps(stageHarness *test_case_harness.TestCaseHarness) error {
	logger := loggerUtils.GetQuietLogger("")

	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	executableDir := b.GetExecutableDirectory()
	gleamTomlFile := filepath.Join(executableDir, "gleam.toml")
	if _, err := os.Stat(gleamTomlFile); os.IsNotExist(err) {
		logger.Criticalf("CodeCrafters internal error. gleam.toml not found at %s", gleamTomlFile)
		logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
		return fmt.Errorf("anti-cheat (ac1) failed")
	}

	gleamToml, err := os.ReadFile(gleamTomlFile)
	if err != nil {
		logger.Criticalf("CodeCrafters internal error. Error reading gleam.toml: %v", err)
		logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
		return fmt.Errorf("anti-cheat (ac1) failed")
	}

	var gleamTomlConfig struct {
		Dependencies []struct {
			Name string
			Path string
		}
	}
	if err := toml.Unmarshal(gleamToml, &gleamTomlConfig); err != nil {
		logger.Criticalf("CodeCrafters internal error. Error unmarshalling gleam.toml: %v", err)
		logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
		return fmt.Errorf("anti-cheat (ac1) failed")
	}

	fmt.Println(gleamTomlConfig.Dependencies)

	for _, dependency := range gleamTomlConfig.Dependencies {
		if dependency.Name == "gleam_stdlib" {
			stdlibPath := filepath.Join(executableDir, dependency.Path)
			if _, err := os.Stat(stdlibPath); os.IsNotExist(err) {
				logger.Criticalf("CodeCrafters internal error. gleam_stdlib not found at %s", stdlibPath)
				logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
				return fmt.Errorf("anti-cheat (ac1) failed")
			}
		}
	}

	return nil
}
