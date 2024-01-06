package operators

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/utils"
)

type NotBlankOperator struct{}

func (NotBlankOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			if fieldType.IsArray() {
				val := data[field]

				if val == nil {
					return false
				}

				if items, err := utils.Array(val); err == nil {
					return len(items) != 0
				}
			} else if fieldType == constants.FieldTypeString {
				val := data[field]

				if val == nil {
					return false
				}

				if str, err := utils.String(val); err == nil {
					return len(str) != 0
				}
			}

			return false
		},
	}
}
