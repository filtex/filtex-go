package operators

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
)

type BlankOperator struct{}

func (BlankOperator) Build(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression {
	if fieldType.IsArray() {
		return &types.MongoExpression{
			Condition: bson.M{
				fmt.Sprintf("%s.0", field): bson.M{
					"$exists": false,
				},
			},
		}
	}

	switch fieldType {
	case constants.FieldTypeString:
		return &types.MongoExpression{
			Condition: bson.M{
				field: bson.M{
					"$exists": true,
					"$eq":     "",
				},
			},
		}
	}

	return nil
}
