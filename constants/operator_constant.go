package constants

import (
	"strings"
)

type Operator struct {
	name  string
	label string
}

func NewOperator(name string, label string) Operator {
	return Operator{
		name:  name,
		label: label,
	}
}

var (
	OperatorUnknown            = NewOperator("", "")
	OperatorEqual              = NewOperator("equal", "Equal")
	OperatorNotEqual           = NewOperator("not-equal", "Not Equal")
	OperatorContain            = NewOperator("contain", "Contain")
	OperatorNotContain         = NewOperator("not-contain", "Not Contain")
	OperatorStartWith          = NewOperator("start-with", "Start With")
	OperatorNotStartWith       = NewOperator("not-start-with", "Not Start With")
	OperatorEndWith            = NewOperator("end-with", "End With")
	OperatorNotEndWith         = NewOperator("not-end-with", "Not End With")
	OperatorBlank              = NewOperator("blank", "Blank")
	OperatorNotBlank           = NewOperator("not-blank", "Not Blank")
	OperatorGreaterThan        = NewOperator("greater-than", "Greater Than")
	OperatorGreaterThanOrEqual = NewOperator("greater-than-or-equal", "Greater Than Or Equal")
	OperatorLessThan           = NewOperator("less-than", "Less Than")
	OperatorLessThanOrEqual    = NewOperator("less-than-or-equal", "Less Than Or Equal")
	OperatorIn                 = NewOperator("in", "In")
	OperatorNotIn              = NewOperator("not-in", "Not In")
)

func (o Operator) String() string {
	return o.name
}

func (o Operator) Equals(str string) bool {
	return strings.ToLower(str) == strings.ToLower(o.name) ||
		strings.ToLower(str) == strings.ToLower(o.label)
}

func ParseOperator(str string) Operator {
	list := []Operator{
		OperatorEqual,
		OperatorNotEqual,
		OperatorContain,
		OperatorNotContain,
		OperatorStartWith,
		OperatorNotStartWith,
		OperatorEndWith,
		OperatorNotEndWith,
		OperatorBlank,
		OperatorNotBlank,
		OperatorGreaterThan,
		OperatorGreaterThanOrEqual,
		OperatorLessThan,
		OperatorLessThanOrEqual,
		OperatorIn,
		OperatorNotIn,
	}

	for _, item := range list {
		if item.Equals(str) {
			return item
		}
	}

	return OperatorUnknown
}
