package test_cases

import (
	"bytes"
	"net/http"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

type GetMoveTestCase struct {
	FEN string
}

func (tc *GetMoveTestCase) Run(stageHarness *test_case_harness.TestCaseHarness) error {
	request, err := http.NewRequest("POST", ADDRESS, bytes.NewBuffer([]byte(`{"fen": "`+tc.FEN+`", "turn": "white", "failed_moves": []}`)))
	if err != nil {
		return err
	}

	test_case := SendRequestTestCase{
		Request:   request,
		Assertion: []assertions.Assertion{&assertions.StatusCodeAssertion{ExpectedStatusCode: 200}, &assertions.ValidMoveAssertion{FEN: tc.FEN}},
	}

	return test_case.Run(stageHarness, stageHarness.Logger)
}
