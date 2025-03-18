package assertions

import (
	"fmt"
	"net/http"
)

type StatusCodeAssertion struct {
	ExpectedStatusCode int
}

func (a *StatusCodeAssertion) Run(response *http.Response) error {
	if response.StatusCode != a.ExpectedStatusCode {
		return fmt.Errorf("expected status code %d, got %d", a.ExpectedStatusCode, response.StatusCode)
	}
	return nil
}
