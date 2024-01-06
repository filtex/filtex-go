package errors

import (
	"errors"
)

var (
	errCouldNotBeBuilt = "could not be built"
)

func NewCouldNotBeBuiltError() error {
	return errors.New(errCouldNotBeBuilt)
}
