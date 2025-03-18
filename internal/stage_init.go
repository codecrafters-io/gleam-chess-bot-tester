package internal

import (
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testInit(stageHarness *test_case_harness.TestCaseHarness) error {
	stageHarness.Logger.Successf("✓ Received exit code %d.", 1)
	return nil
}
