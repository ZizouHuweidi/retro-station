package errors

import (
	customErrors "github.com/zizouhuweidi/retro-station/internal/pkg/http/http_errors/custom_errors"

	"emperror.dev/errors"
)

type readStreamError struct {
	customErrors.InternalServerError
}
type ReadStreamError interface {
	customErrors.InternalServerError
	IsReadStreamError() bool
}

func NewReadStreamError(err error) error {
	internal := customErrors.NewInternalServerErrorWrap(err, "unable to read events from stream")
	customErr := customErrors.GetCustomError(internal)

	br := &readStreamError{
		InternalServerError: customErr.(customErrors.InternalServerError),
	}

	return errors.WithStackIf(br)
}

func (err *readStreamError) IsReadStreamError() bool {
	return true
}

func IsReadStreamError(err error) bool {
	var rs ReadStreamError
	if errors.As(err, &rs) {
		return rs.IsReadStreamError()
	}

	return false
}
