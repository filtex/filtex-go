package errors

import (
	"errors"
)

var (
	errInvalidFieldType  = "invalid field type"
	errInvalidFieldName  = "invalid field name"
	errInvalidFieldLabel = "invalid field label"
)

func NewInvalidFieldTypeError() error {
	return errors.New(errInvalidFieldType)
}

func NewInvalidFieldNameError() error {
	return errors.New(errInvalidFieldName)
}

func NewInvalidFieldLabelError() error {
	return errors.New(errInvalidFieldLabel)
}
