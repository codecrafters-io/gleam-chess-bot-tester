package assertions

import "net/http"

type Assertion interface {
	Run(response *http.Response) error
}
