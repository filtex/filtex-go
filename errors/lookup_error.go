package errors

import (
	"errors"
)

var (
	errInvalidLookupKey    = "invalid field key"
	errInvalidLookupValues = "invalid field values"
)

func NewInvalidLookupKeyError() error {
	return errors.New(errInvalidLookupKey)
}

func NewInvalidLookupValuesError() error {
	return errors.New(errInvalidLookupValues)
}
