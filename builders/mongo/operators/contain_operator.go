package operators

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type ContainOperator struct{}

func (ContainOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() {
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$in": []interface{}{value},
				},
			},
		}
	}

	switch fieldType {
	case constants.FieldTypeString:
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$regex":   value,
					"$options": "i",
				},
			},
		}
	}

	return nil
}
