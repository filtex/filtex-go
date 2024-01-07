package operators

import (
	"fmt"
	"strings"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/utils"
)

type InOperator struct{}

func (InOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType.IsArray() || value == nil {
		return nil
	}

	if !utils.IsArray(value) {
		if fieldType == constants.FieldTypeString {
			return &types.PostgresExpression{
				Condition: fmt.Sprintf("LOWER(%s) IN (LOWER($%v))", field, index),
				Args:      []interface{}{value},
			}
		} else {
			return &types.PostgresExpression{
				Condition: fmt.Sprintf("%s IN ($%v)", field, index),
				Args:      []interface{}{value},
			}
		}
	}

	if fieldType == constants.FieldTypeString {
		indexes := make([]string, 0)

		for i := index; i < index+len(value.([]interface{})); i++ {
			indexes = append(indexes, fmt.Sprintf("LOWER($%v)", i))
		}

		return &types.PostgresExpression{
			Condition: fmt.Sprintf("LOWER(%s) IN (%s)", field, strings.Join(indexes, ",")),
			Args:      value.([]interface{}),
		}
	} else {
		indexes := make([]string, 0)

		for i := index; i < index+len(value.([]interface{})); i++ {
			indexes = append(indexes, fmt.Sprintf("$%v", i))
		}

		return &types.PostgresExpression{
			Condition: fmt.Sprintf("%s IN (%s)", field, strings.Join(indexes, ",")),
			Args:      value.([]interface{}),
		}
	}
}
