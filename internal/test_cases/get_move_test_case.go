package test_cases

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

type GetMoveTestCase struct {
	FEN                        string
	AssertGeneratedMoveIsValid bool
}

// For invalid FENs, we still need to parse the turn
// Parsing FENs is a no go for those test cases
func getTurn(fenStr string) string {
	parts := strings.Split(fenStr, " ")
	return strings.TrimSpace(parts[1])
}

func (tc *GetMoveTestCase) Run(stageHarness *test_case_harness.TestCaseHarness) error {
	request, err := http.NewRequest("POST", ADDRESS, bytes.NewBuffer([]byte(`{"fen": "`+tc.FEN+`", "turn": "`+getTurn(tc.FEN)+`", "failed_moves": []}`)))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	allAssertions := []assertions.Assertion{&assertions.StatusCodeAssertion{ExpectedStatusCode: 200}}
	if tc.AssertGeneratedMoveIsValid {
		allAssertions = append(allAssertions, &assertions.ValidMoveAssertion{FEN: tc.FEN})
	}

	test_case := SendRequestTestCase{
		Request:   request,
		Assertion: allAssertions,
	}

	return test_case.Run(stageHarness, stageHarness.Logger)
}
