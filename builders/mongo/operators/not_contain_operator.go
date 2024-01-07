package operators

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type NotContainOperator struct{}

func (NotContainOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() {
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$nin": []interface{}{value},
				},
			},
		}
	}

	switch fieldType {
	case constants.FieldTypeString:
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$regex":   fmt.Sprintf("^((?!%s).)*$", value),
					"$options": "i",
				},
			},
		}
	}

	return nil
}
