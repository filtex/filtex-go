package operators

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/utils"
)

type LessThanOperator struct{}

func (LessThanOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			val := data[field]

			if val == nil {
				return false
			}

			switch fieldType {
			case constants.FieldTypeNumber:
				castedResultValue, castedResultValueErr := utils.Number(val)
				castedValue, castedValueErr := utils.Number(value)
				if castedResultValueErr == nil && castedValueErr == nil {
					return castedResultValue < castedValue
				}
			case constants.FieldTypeDate:
				castedResultValue, castedResultValueErr := utils.Date(val)
				castedValue, castedValueErr := utils.Date(value)
				if castedResultValueErr == nil && castedValueErr == nil {
					return castedResultValue.UnixNano() < castedValue.UnixNano()
				}
			case constants.FieldTypeTime:
				castedResultValue, castedResultValueErr := utils.Time(val)
				castedValue, castedValueErr := utils.Time(value)
				if castedResultValueErr == nil && castedValueErr == nil {
					return *castedResultValue < *castedValue
				}
			case constants.FieldTypeDateTime:
				castedResultValue, castedResultValueErr := utils.DateTime(val)
				castedValue, castedValueErr := utils.DateTime(value)
				if castedResultValueErr == nil && castedValueErr == nil {
					return castedResultValue.UnixNano() < castedValue.UnixNano()
				}
			}

			return false
		},
	}
}
