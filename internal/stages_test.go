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
			UntilStageSlug:      "gd8",
			CodePath:            "./test_helpers/scenarios/test_bot",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/test_bot/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"invalid_move": {
			UntilStageSlug:      "zz5",
			CodePath:            "./test_helpers/scenarios/failure_bot",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/test_bot/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"invalid_repo_state": {
			StageSlugs:          []string{"si0"},
			CodePath:            "./test_helpers/scenarios/failure_bot",
			ExpectedExitCode:    1,
			StdoutFixturePath:   "./test_helpers/fixtures/invalid_repo_state/failure",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
	}

	tester_utils_testing.TestTesterOutput(t, testerDefinition, testCases)
}

func normalizeTesterOutput(testerOutput []byte) []byte {
	return testerOutput
}
