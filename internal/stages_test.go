package internal

import (
	"testing"

	tester_utils_testing "github.com/codecrafters-io/tester-utils/testing"
)

func TestStages(t *testing.T) {
	testCases := map[string]tester_utils_testing.TesterOutputTestCase{
		"success": {
			UntilStageSlug:      "a05",
			CodePath:            "./test_helpers/scenarios/test_bot",
			ExpectedExitCode:    0,
			StdoutFixturePath:   "./test_helpers/fixtures/test_bot/success",
			NormalizeOutputFunc: normalizeTesterOutput,
		},
		"invalid_move": {
			UntilStageSlug:      "a05",
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
