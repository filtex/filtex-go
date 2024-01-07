package operators

import (
	"fmt"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
)

type NotEqualOperator struct{}

func (NotEqualOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType.IsArray() {
		return nil
	}

	if fieldType == constants.FieldTypeString {
		return &types.PostgresExpression{
			Condition: fmt.Sprintf("%s NOT ILIKE $%v", field, index),
			Args:      []interface{}{value},
		}
	}

	return &types.PostgresExpression{
		Condition: fmt.Sprintf("%s <> $%v", field, index),
		Args:      []interface{}{value},
	}
}
