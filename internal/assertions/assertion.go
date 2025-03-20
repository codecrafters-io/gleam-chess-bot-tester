package assertions

import (
	"net/http"

	"github.com/codecrafters-io/tester-utils/logger"
)

type Assertion interface {
	Run(response *http.Response, logger *logger.Logger) error
}
