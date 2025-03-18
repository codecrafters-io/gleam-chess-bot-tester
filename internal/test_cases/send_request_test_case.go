package test_cases

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

const PROTOCOL = "http"
const HOST = "localhost"
const PORT = "8000"
const ENDPOINT = "/move"
const ADDRESS = PROTOCOL + "://" + HOST + ":" + PORT + ENDPOINT

type SendRequestTestCase struct {
	Request   *http.Request
	Assertion []assertions.Assertion
}

func (tc *SendRequestTestCase) Run(harness *test_case_harness.TestCaseHarness, logger *logger.Logger) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// sleep for 1 second
	// TODO: Remove this once we've confirmed that the server is running
	time.Sleep(1 * time.Second)

	// Make the request
	logger.Debugf("Making HTTP request to %s", tc.Request.URL)
	response, err := client.Do(tc.Request)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer response.Body.Close()

	// Run the assertion
	if tc.Assertion != nil {
		for _, assertion := range tc.Assertion {
			if err := assertion.Run(response, logger); err != nil {
				return fmt.Errorf("assertion failed: %v", err)
			}
		}
	}

	logger.Successf("Request completed successfully")
	return nil
}
