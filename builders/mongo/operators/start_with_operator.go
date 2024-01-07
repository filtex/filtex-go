package operators

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type StartWithOperator struct{}

func (StartWithOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType != constants.FieldTypeString {
		return nil
	}

	return &types.MongoExpression{
		Condition: bson.M{
			field: bson.M{
				"$regex":   fmt.Sprintf("^%s", value),
				"$options": "i",
			},
		},
	}
}
