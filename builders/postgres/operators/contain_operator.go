package operators

import (
	"fmt"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
)

type ContainOperator struct{}

func (ContainOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType.IsArray() {
		if fieldType == constants.FieldTypeStringArray {
			return &types.PostgresExpression{
				Condition: fmt.Sprintf("LOWER($%v) = ANY (LOWER(%s::TEXT)::TEXT[])", index, field),
				Args:      []interface{}{value},
			}
		}

		return &types.PostgresExpression{
			Condition: fmt.Sprintf("$%v = ANY (%s)", index, field),
			Args:      []interface{}{value},
		}
	}

	switch fieldType {
	case constants.FieldTypeString:
		return &types.PostgresExpression{
			Condition: fmt.Sprintf("%s ILIKE '%%' || $%v || '%%'", field, index),
			Args:      []interface{}{value},
		}
	}

	return nil
}
