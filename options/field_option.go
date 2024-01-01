package options

import (
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/models"
)

type FieldOption struct {
	name       string
	label      string
	lookup     string
	fieldType  constants.FieldType
	isArray    bool
	isNullable bool
}

func NewFieldOption() *FieldOption {
	return &FieldOption{}
}

func (f *FieldOption) String() *FieldOption {
	f.fieldType = constants.FieldTypeString
	return f
}

func (f *FieldOption) Number() *FieldOption {
	f.fieldType = constants.FieldTypeNumber
	return f
}

func (f *FieldOption) Boolean() *FieldOption {
	f.fieldType = constants.FieldTypeBoolean
	return f
}

func (f *FieldOption) Date() *FieldOption {
	f.fieldType = constants.FieldTypeDate
	return f
}

func (f *FieldOption) Time() *FieldOption {
	f.fieldType = constants.FieldTypeTime
	return f
}

func (f *FieldOption) DateTime() *FieldOption {
	f.fieldType = constants.FieldTypeDateTime
	return f
}

func (f *FieldOption) Array() *FieldOption {
	f.isArray = true
	return f
}

func (f *FieldOption) Nullable() *FieldOption {
	f.isNullable = true
	return f
}

func (f *FieldOption) Name(name string) *FieldOption {
	f.name = name
	return f
}

func (f *FieldOption) Label(label string) *FieldOption {
	f.label = label
	return f
}

func (f *FieldOption) Lookup(lookup string) *FieldOption {
	f.lookup = lookup
	return f
}

func (f *FieldOption) Build(lookups map[string][]models.Lookup) (*models.Field, error) {
	if f.fieldType == "" {
		return nil, errors.NewInvalidFieldTypeError()
	}

	if f.name == "" {
		return nil, errors.NewInvalidFieldNameError()
	}

	if f.label == "" {
		return nil, errors.NewInvalidFieldLabelError()
	}

	fieldType := f.fieldType

	if f.isArray {
		if fieldType == constants.FieldTypeString {
			fieldType = constants.FieldTypeStringArray
		} else if fieldType == constants.FieldTypeNumber {
			fieldType = constants.FieldTypeNumberArray
		} else if fieldType == constants.FieldTypeBoolean {
			fieldType = constants.FieldTypeBooleanArray
		} else if fieldType == constants.FieldTypeDate {
			fieldType = constants.FieldTypeDateArray
		} else if fieldType == constants.FieldTypeTime {
			fieldType = constants.FieldTypeTimeArray
		} else if fieldType == constants.FieldTypeDateTime {
			fieldType = constants.FieldTypeDateTimeArray
		}
	}

	fieldValues := make([]models.Lookup, 0)

	for k, v := range lookups {
		if k == f.lookup {
			fieldValues = v
		}
	}

	operators := make([]string, 0)

	if !f.isArray {
		operators = append(operators, constants.OperatorEqual.String())
		operators = append(operators, constants.OperatorNotEqual.String())
	}

	if (fieldType == constants.FieldTypeNumber ||
		fieldType == constants.FieldTypeNumberArray ||
		fieldType == constants.FieldTypeDate ||
		fieldType == constants.FieldTypeDateArray ||
		fieldType == constants.FieldTypeTime ||
		fieldType == constants.FieldTypeTimeArray ||
		fieldType == constants.FieldTypeDateTime ||
		fieldType == constants.FieldTypeDateTimeArray) && len(fieldValues) == 0 {
		operators = append(operators, constants.OperatorGreaterThan.String())
		operators = append(operators, constants.OperatorGreaterThanOrEqual.String())
		operators = append(operators, constants.OperatorLessThan.String())
		operators = append(operators, constants.OperatorLessThanOrEqual.String())
	}

	if f.isArray || f.isNullable {
		operators = append(operators, constants.OperatorBlank.String())
		operators = append(operators, constants.OperatorNotBlank.String())
	}

	if f.isArray {
		operators = append(operators, constants.OperatorContain.String())
		operators = append(operators, constants.OperatorNotContain.String())
	} else if fieldType == constants.FieldTypeString && len(fieldValues) == 0 {
		operators = append(operators, constants.OperatorContain.String())
		operators = append(operators, constants.OperatorNotContain.String())
		operators = append(operators, constants.OperatorStartWith.String())
		operators = append(operators, constants.OperatorNotStartWith.String())
		operators = append(operators, constants.OperatorEndWith.String())
		operators = append(operators, constants.OperatorNotEndWith.String())
	}

	if !f.isArray {
		operators = append(operators, constants.OperatorIn.String())
		operators = append(operators, constants.OperatorNotIn.String())
	}

	return &models.Field{
		Name:      f.name,
		Type:      fieldType.String(),
		Label:     f.label,
		Operators: operators,
		Values:    fieldValues,
	}, nil
}
