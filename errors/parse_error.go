package errors

import (
	"errors"
)

var (
	errOperatorCouldNotBeParsed = "invalid operator"
	errLogicCouldNotBeParsed    = "invalid logic"
	errCouldNotBeParsed         = "could not be parsed"
)

func NewOperatorCouldNotBeParsedError() error {
	return errors.New(errOperatorCouldNotBeParsed)
}

func NewLogicCouldNotBeParsedError() error {
	return errors.New(errLogicCouldNotBeParsed)
}

func NewCouldNotBeParsedError() error {
	return errors.New(errCouldNotBeParsed)
}
