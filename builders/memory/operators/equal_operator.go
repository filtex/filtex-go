package operators

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/builders/memory/utils"
	"github.com/filtex/filtex-go/constants"
)

type EqualOperator struct{}

func (EqualOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			if fieldType.IsArray() {
				return false
			}

			val := data[field]

			if val == nil {
				return false
			}

			return utils.CheckEquality(fieldType, val, value)
		},
	}
}
