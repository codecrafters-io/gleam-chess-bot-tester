package assertions

import (
	"fmt"
	"net/http"

	"github.com/codecrafters-io/tester-utils/logger"
)

type StatusCodeAssertion struct {
	ExpectedStatusCode int
}

func (a *StatusCodeAssertion) Run(response *http.Response, logger *logger.Logger) error {
	if response.StatusCode != a.ExpectedStatusCode {
		return fmt.Errorf("Expected status code %d, got %d", a.ExpectedStatusCode, response.StatusCode)
	}

	logger.Successf("✓ Received status code %d", a.ExpectedStatusCode)
	return nil
}
