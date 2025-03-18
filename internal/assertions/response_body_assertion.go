package assertions

import (
	"fmt"
	"io"
	"net/http"

	"github.com/codecrafters-io/tester-utils/logger"
)

type ResponseBodyAssertion struct {
	ExpectedBody string
}

func (a *ResponseBodyAssertion) Run(response *http.Response, logger *logger.Logger) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %v", err)
	}

	if string(body) != a.ExpectedBody {
		return fmt.Errorf("Expected response body %s, got %s", a.ExpectedBody, string(body))
	}

	logger.Successf("✓ Received body %s", a.ExpectedBody)
	return nil
}
