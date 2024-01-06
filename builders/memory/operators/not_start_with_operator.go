package operators

import (
	"strings"

	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/constants"
)

type NotStartWithOperator struct{}

func (NotStartWithOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			val := data[field]

			if val == nil {
				return true
			}

			if fieldType == constants.FieldTypeString {
				return !strings.HasPrefix(strings.ToLower(val.(string)), strings.ToLower(value.(string)))
			}

			return false
		},
	}
}
