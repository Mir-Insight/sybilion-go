package devportalclient

import (
	"errors"
	"fmt"

	devportal "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

// APIError wraps a non-2xx API response when the body could not be decoded into a typed model.
type APIError struct {
	StatusCode int
	Body       []byte
	Message    string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("api error: status=%d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("api error: status=%d", e.StatusCode)
}

// AsGenericOpenAPIError unwraps openapi-generator's GenericOpenAPIError when possible.
func AsGenericOpenAPIError(err error) (*devportal.GenericOpenAPIError, bool) {
	var ge *devportal.GenericOpenAPIError
	if errors.As(err, &ge) {
		return ge, true
	}
	return nil, false
}
