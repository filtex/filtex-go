package operators

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	memoryUtils "github.com/filtex/filtex-go/builders/memory/utils"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/utils"
	"strings"
)

type NotContainOperator struct{}

func (NotContainOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			if fieldType.IsArray() {
				val := data[field]

				if val == nil {
					return true
				}

				if items, err := utils.Array(val); err == nil {
					for _, item := range items {
						if memoryUtils.CheckEquality(fieldType, item, value) {
							return false
						}
					}

					return true
				}
			} else if fieldType == constants.FieldTypeString {
				val := data[field]

				if val == nil {
					return true
				}

				return !strings.Contains(strings.ToLower(val.(string)), strings.ToLower(value.(string)))
			}

			return false
		},
	}
}
