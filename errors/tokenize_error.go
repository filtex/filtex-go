package errors

import (
	"errors"
)

var (
	errCouldNotBeTokenized = "could not be tokenized"
)

func NewCouldNotBeTokenizedError() error {
	return errors.New(errCouldNotBeTokenized)
}
