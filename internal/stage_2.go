package internal

import (
	"net/http"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	chess_bot_executable "github.com/codecrafters-io/gleam-chess-bot-tester/internal/chess-bot-executable"
	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func test2(stageHarness *test_case_harness.TestCaseHarness) error {
	b := chess_bot_executable.NewChessBotExecutable(stageHarness)
	if err := b.Run(); err != nil {
		return err
	}

	logger := stageHarness.Logger

	request, err := http.NewRequest("GET", test_cases.ADDRESS, nil)
	if err != nil {
		return err
	}

	test_case := test_cases.SendRequestTestCase{
		Request:   request,
		Assertion: []assertions.Assertion{&assertions.StatusCodeAssertion{ExpectedStatusCode: 200}, &assertions.ResponseBodyAssertion{ExpectedBody: "expected response here"}},
	}

	return test_case.Run(stageHarness, logger)
}
