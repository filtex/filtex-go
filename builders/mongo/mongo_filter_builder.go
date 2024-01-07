package mongo

import (
	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
)

type MongoFilterBuilder struct {
}

func NewMongoFilterBuilder() *MongoFilterBuilder {
	return &MongoFilterBuilder{}
}

func (b *MongoFilterBuilder) Build(expression expressions.Expression) (*types.MongoExpression, error) {
	return nil, errors.NewCouldNotBeBuiltError()
}
