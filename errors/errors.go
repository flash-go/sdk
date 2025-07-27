package errors

import (
	"errors"
	"fmt"
)

type Error error

var (
	ErrBadRequest         Error = errors.New("bad_request")
	ErrUnauthorized       Error = errors.New("unauthorized")
	ErrForbidden          Error = errors.New("forbidden")
	ErrServiceUnavailable Error = errors.New("service_unavailable")
	ErrNotFound           Error = errors.New("not_found")
)

func New(base Error, message string) error {
	return fmt.Errorf("%w:%s", base, message)
}
