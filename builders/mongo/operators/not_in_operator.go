package operators

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type NotInOperator struct{}

func (NotInOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() || value == nil {
		return nil
	}

	items, ok := value.([]interface{})
	if !ok {
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$nin": []interface{}{
						value,
					},
				},
			},
		}
	}

	return &types.MongoExpression{
		Condition: bson.M{
			field: bson.M{
				"$nin": items,
			},
		},
	}
}
