package test_cases

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codecrafters-io/gleam-chess-bot-tester/internal/assertions"
	"github.com/codecrafters-io/tester-utils/logger"
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

func (tc *SendRequestTestCase) Run(logger *logger.Logger) error {
	client := &http.Client{}

	// sleep for 1 second
	// TODO: Remove this once we've confirmed that the server is running
	time.Sleep(1 * time.Second)

	startTime := time.Now()

	// Make the request
	logger.Debugf("Making HTTP request to %s", tc.Request.URL)
	response, err := client.Do(tc.Request)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer response.Body.Close()

	// Create Response struct set duration on it
	// Use that for all assertions
	duration := time.Since(startTime)
	if duration > 5*time.Second {
		return fmt.Errorf("request took too long to complete: %s", duration)
	}
	// logger.Debugf("Request completed in %s", duration)

	// Run the assertion
	if tc.Assertion != nil {
		for _, assertion := range tc.Assertion {
			if err := assertion.Run(response, logger); err != nil {
				return fmt.Errorf("Assertion failed: %v", err)
			}
		}
	}

	logger.Successf("Request completed successfully")
	return nil
}
