package operators

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type NotEqualOperator struct{}

func (NotEqualOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() {
		return nil
	}

	if fieldType == constants.FieldTypeString {
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$not": bson.M{
						"$regex":   fmt.Sprintf("^%s$", value),
						"$options": "i",
					},
				},
			},
		}
	}

	return &types.MongoExpression{
		Condition: bson.M{
			field: bson.M{
				"$ne": value,
			},
		},
	}
}
