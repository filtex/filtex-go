package operators

import (
	"fmt"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
)

type GreaterThanOrEqualOperator struct{}

func (GreaterThanOrEqualOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType != constants.FieldTypeNumber &&
		fieldType != constants.FieldTypeDate &&
		fieldType != constants.FieldTypeTime &&
		fieldType != constants.FieldTypeDateTime {
		return nil
	}

	return &types.PostgresExpression{
		Condition: fmt.Sprintf("%s >= $%v", field, index),
		Args:      []interface{}{value},
	}
}
