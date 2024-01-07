package operators

import (
	"fmt"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
)

type NotBlankOperator struct{}

func (NotBlankOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType.IsArray() {
		return &types.PostgresExpression{
			Condition: fmt.Sprintf("ARRAY_LENGTH(%s, 1) <> 0", field),
			Args:      []interface{}{},
		}
	}

	switch fieldType {
	case constants.FieldTypeString:
		return &types.PostgresExpression{
			Condition: fmt.Sprintf("%s IS NOT NULL AND %s <> ''", field, field),
			Args:      []interface{}{},
		}
	}

	return nil
}
