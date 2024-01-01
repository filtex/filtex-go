package errors

import (
	"errors"
)

var (
	errInvalidField        = "invalid field"
	errInvalidOperator     = "invalid operator"
	errInvalidValue        = "invalid value"
	errInvalidLogic        = "invalid logic"
	errInvalidToken        = "invalid token"
	errInvalidLastToken    = "invalid last token"
	errMismatchedBrackets  = "mismatched brackets"
	errCouldNotBeValidated = "could not be validated"
)

func NewInvalidFieldError() error {
	return errors.New(errInvalidField)
}

func NewInvalidOperatorError() error {
	return errors.New(errInvalidOperator)
}

func NewInvalidValueError() error {
	return errors.New(errInvalidValue)
}

func NewInvalidLogicError() error {
	return errors.New(errInvalidLogic)
}

func NewInvalidTokenError() error {
	return errors.New(errInvalidToken)
}

func NewInvalidLastTokenError() error {
	return errors.New(errInvalidLastToken)
}

func NewMismatchedBracketsError() error {
	return errors.New(errMismatchedBrackets)
}

func NewCouldNotBeValidatedError() error {
	return errors.New(errCouldNotBeValidated)
}
