package operators

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type GreaterThanOrEqualOperator struct{}

func (GreaterThanOrEqualOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType != constants.FieldTypeNumber &&
		fieldType != constants.FieldTypeDate &&
		fieldType != constants.FieldTypeTime &&
		fieldType != constants.FieldTypeDateTime {
		return nil
	}

	return &types.MongoExpression{
		Condition: bson.M{
			field: bson.M{
				"$gte": value,
			},
		},
	}
}
