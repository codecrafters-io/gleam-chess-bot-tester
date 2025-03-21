package internal

import (
	"os"
	"testing"

	tester_utils_testing "github.com/codecrafters-io/tester-utils/testing"
)

func TestStages(t *testing.T) {
	os.Setenv("CODECRAFTERS_RANDOM_SEED", "1234567890")

	testCases := map[string]tester_utils_testing.TesterOutputTestCase{
		"success": {
			UntilStageSlug:      "a04",
			CodePath:            "./test_helpers/scenarios/test_bot",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/test_bot/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"success_gleam": {
			UntilStageSlug:      "a02",
			CodePath:            "./test_helpers/scenarios/gleam",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/gleam/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"invalid_move": {
			UntilStageSlug:      "a04",
			CodePath:            "./test_helpers/scenarios/failure_bot",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/test_bot/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
