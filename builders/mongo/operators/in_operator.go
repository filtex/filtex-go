package operators

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type InOperator struct{}

func (InOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() || value == nil {
		return nil
	}

	items, ok := value.([]interface{})
	if !ok {
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$in": []interface{}{
						value,
					},
				},
			},
		}
	}

	return &types.MongoExpression{
		Condition: bson.M{
			field: bson.M{
				"$in": items,
			},
		},
	}
}
