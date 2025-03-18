package test_cases

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

const HOST = "localhost"
const PORT = "8000"
const ENDPOINT = "/move"
const ADDRESS = HOST + ":" + PORT + ENDPOINT

type SendRequestTestCase struct {
	Request   *http.Request
	Assertion []func(*http.Response) error
}

func (tc *SendRequestTestCase) Run(harness *test_case_harness.TestCaseHarness, logger *logger.Logger) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// // Wait a bit for the server to start
	// time.Sleep(1 * time.Second)

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
			if err := assertion(response); err != nil {
				return fmt.Errorf("assertion failed: %v", err)
			}
		}
	}

	logger.Successf("Request completed successfully")
	return nil
}

func StatusCodeIs200(response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got %d", response.StatusCode)
	}

	return nil
}

func BodyIs(response *http.Response, expectedBody string) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if string(body) != expectedBody {
		return fmt.Errorf("expected body %s, got %s", expectedBody, string(body))
	}

	return nil
}
