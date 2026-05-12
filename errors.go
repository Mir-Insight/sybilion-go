package sybilion

import (
	"errors"
	"fmt"

	api "go.sybilion.dev/sybilion/api"
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
func AsGenericOpenAPIError(err error) (*api.GenericOpenAPIError, bool) {
	var ge *api.GenericOpenAPIError
	if errors.As(err, &ge) {
		return ge, true
	}
	return nil, false
}
