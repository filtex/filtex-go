package mongo

import (
	"github.com/filtex/filtex-go/builders/mongo/logics"
	"github.com/filtex/filtex-go/builders/mongo/operators"
	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
)

type MongoFilterBuilder struct {
	logicsMap    map[constants.Logic]func(expressions []*types.MongoExpression) *types.MongoExpression
	operatorsMap map[constants.Operator]func(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression
}

func NewMongoFilterBuilder() *MongoFilterBuilder {
	return &MongoFilterBuilder{
		logicsMap: map[constants.Logic]func(expressions []*types.MongoExpression) *types.MongoExpression{
			constants.LogicAnd: logics.AndLogic{}.Build,
			constants.LogicOr:  logics.OrLogic{}.Build,
		},
		operatorsMap: map[constants.Operator]func(fieldType constants.FieldType, field string, value interface{}) *types.MongoExpression{
			constants.OperatorEqual:              operators.EqualOperator{}.Build,
			constants.OperatorNotEqual:           operators.NotEqualOperator{}.Build,
			constants.OperatorContain:            operators.ContainOperator{}.Build,
			constants.OperatorNotContain:         operators.NotContainOperator{}.Build,
			constants.OperatorStartWith:          operators.StartWithOperator{}.Build,
			constants.OperatorNotStartWith:       operators.NotStartWithOperator{}.Build,
			constants.OperatorEndWith:            operators.EndWithOperator{}.Build,
			constants.OperatorNotEndWith:         operators.NotEndWithOperator{}.Build,
			constants.OperatorBlank:              operators.BlankOperator{}.Build,
			constants.OperatorNotBlank:           operators.NotBlankOperator{}.Build,
			constants.OperatorGreaterThan:        operators.GreaterThanOperator{}.Build,
			constants.OperatorGreaterThanOrEqual: operators.GreaterThanOrEqualOperator{}.Build,
			constants.OperatorLessThan:           operators.LessThanOperator{}.Build,
			constants.OperatorLessThanOrEqual:    operators.LessThanOrEqualOperator{}.Build,
			constants.OperatorIn:                 operators.InOperator{}.Build,
			constants.OperatorNotIn:              operators.NotInOperator{}.Build,
		},
	}
}

func (b *MongoFilterBuilder) Build(expression expressions.Expression) (*types.MongoExpression, error) {
	switch exp := expression.(type) {
	case *expressions.LogicExpression:
		expressions := make([]*types.MongoExpression, 0)

		for _, v := range exp.Expressions {
			e, err := b.Build(v)
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, e)
		}

		if fn, ok := b.logicsMap[exp.Logic]; ok {
			return fn(expressions), nil
		}

		return nil, errors.NewCouldNotBeBuiltError()
	case *expressions.OperatorExpression:
		if fn, ok := b.operatorsMap[exp.Operator]; ok {
			return fn(exp.Type, exp.Field, exp.Value), nil
		}

		return nil, errors.NewCouldNotBeBuiltError()
	}

	return nil, errors.NewCouldNotBeBuiltError()
}
