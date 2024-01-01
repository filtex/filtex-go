package errors

import (
	"errors"
)

var (
	errCouldNotBeCasted = "could not be casted"
)

func NewCouldNotBeCastedError() error {
	return errors.New(errCouldNotBeCasted)
}
