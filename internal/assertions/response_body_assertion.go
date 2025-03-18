package assertions

import (
	"fmt"
	"io"
	"net/http"
)

type ResponseBodyAssertion struct {
	ExpectedBody string
}

func (a *ResponseBodyAssertion) Run(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if string(body) != a.ExpectedBody {
		return fmt.Errorf("expected body %s, got %s", a.ExpectedBody, string(body))
	}

	return nil
}
