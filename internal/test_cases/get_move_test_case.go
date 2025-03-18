package test_cases

import (
	"bytes"
	"net/http"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
	"github.com/corentings/chess"
)

type GetMoveTestCase struct {
	FEN                        string
	AssertGeneratedMoveIsValid bool
}

func getTurn(fenStr string) string {
	fen, err := chess.FEN(fenStr)
	if err != nil {
		panic("Failed to parse FEN: " + err.Error())
	}

	game := chess.NewGame(fen)
	return game.Position().Turn().String()
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
