package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/BurntSushi/toml"
	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	loggerUtils "github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var WHITELISTED_DEPENDENCIES = []string{
	"gleam_stdlib",
	"gleam_http",
	"gleam_json",
	"gleam_time",
	"gleam_community_maths",
	"flash",
	"iv",
	"glearray",
	"snag",
	"birl",
	"gtempo",
	"gleam_erlang",
	"gleam_otp",
	"mist",
	"wisp",
	"gleam_javascript",
	"glen",
}

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
		Dependencies map[string]string `toml:"dependencies"`
	}
	if err := toml.Unmarshal(gleamToml, &gleamTomlConfig); err != nil {
		logger.Criticalf("CodeCrafters internal error. Error unmarshalling gleam.toml: %v", err)
		logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
		return fmt.Errorf("anti-cheat (ac2) failed")
	}

	deps_3p := make([]string, 0, len(gleamTomlConfig.Dependencies))
	for k := range gleamTomlConfig.Dependencies {
		deps_3p = append(deps_3p, k)
	}

	for _, dep := range deps_3p {
		if !slices.Contains(WHITELISTED_DEPENDENCIES, dep) {
			logger.Criticalf("CodeCrafters internal error. Dependency %s is not allowed", dep)
			logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
			return fmt.Errorf("anti-cheat (ac1) failed")
		}
	}

	return nil
}

func antiCheatFFI(stageHarness *test_case_harness.TestCaseHarness) error {
	logger := loggerUtils.GetQuietLogger("")

	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	executableDir := b.GetExecutableDirectory()

	var allFiles []string
	err := filepath.Walk(executableDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			allFiles = append(allFiles, path)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	for _, file := range allFiles {
		if strings.HasSuffix(file, ".gleam") {
			content, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("error reading file: %v", err)
			}
			if strings.Contains(string(content), "@external") {
				logger.Criticalf("CodeCrafters internal error. External FFI functions are not allowed.\nFile: %s", file)
				logger.Criticalf("Try again? Please contact us at hello@codecrafters.io if this persists.")
				return fmt.Errorf("anti-cheat (ac3) failed")
			}
		}
	}

	return nil
}
