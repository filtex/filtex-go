package operators

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/builders/memory/utils"
	"github.com/filtex/filtex-go/constants"
)

type InOperator struct{}

func (InOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			val := data[field]

			if fieldType.IsArray() || value == nil {
				return false
			}

			items, ok := value.([]interface{})
			if !ok {
				return utils.CheckEquality(fieldType, val, value)
			}

			for _, v := range items {
				if utils.CheckEquality(fieldType, val, v) {
					return true
				}
			}

			return false
		},
	}
}
