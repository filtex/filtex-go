package expressions

import (
	"github.com/filtex/filtex-go/constants"
)

type OperatorExpression struct {
	Type     constants.FieldType
	Field    string
	Operator constants.Operator
	Value    interface{}
}

func NewOperatorExpression(fieldType constants.FieldType, field string, operator constants.Operator, value interface{}) Expression {
	return &OperatorExpression{
		Type:     fieldType,
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}
