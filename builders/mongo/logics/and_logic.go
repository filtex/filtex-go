package logics

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/filtex/filtex-go/builders/mongo/types"
)

type AndLogic struct{}

func (AndLogic) Build(expressions []*types.MongoExpression) *types.MongoExpression {
	conditions := make([]bson.M, 0)

	for _, v := range expressions {
		conditions = append(conditions, v.Condition)
	}

	return &types.MongoExpression{
		Condition: bson.M{
			"$and": conditions,
		},
	}
}
