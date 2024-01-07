package postgres

import (
	"github.com/filtex/filtex-go/builders/postgres/logics"
	"github.com/filtex/filtex-go/builders/postgres/operators"
	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
)

type PostgresFilterBuilder struct {
	logicsMap    map[constants.Logic]func(expressions []types.PostgresExpression) *types.PostgresExpression
	operatorsMap map[constants.Operator]func(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression
}

func NewPostgresFilterBuilder() *PostgresFilterBuilder {
	return &PostgresFilterBuilder{
		logicsMap: map[constants.Logic]func(expressions []types.PostgresExpression) *types.PostgresExpression{
			constants.LogicAnd: logics.AndLogic{}.Build,
			constants.LogicOr:  logics.OrLogic{}.Build,
		},
		operatorsMap: map[constants.Operator]func(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression{
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

func (b *PostgresFilterBuilder) Build(ex expressions.Expression) (*types.PostgresExpression, error) {
	index := 1
	return b.buildInternal(ex, &index)
}

func (b *PostgresFilterBuilder) buildInternal(ex expressions.Expression, index *int) (*types.PostgresExpression, error) {
	switch exp := ex.(type) {
	case *expressions.LogicExpression:
		expressions := make([]types.PostgresExpression, 0)

		for _, v := range exp.Expressions {
			e, err := b.buildInternal(v, index)
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, *e)
		}

		if fn, ok := b.logicsMap[exp.Logic]; ok {
			return fn(expressions), nil
		}

		return nil, errors.NewCouldNotBeBuiltError()
	case *expressions.OperatorExpression:
		if fn, ok := b.operatorsMap[exp.Operator]; ok {
			result := fn(exp.Type, exp.Field, exp.Value, *index)
			*index += len(result.Args)
			return result, nil
		}

		return nil, errors.NewCouldNotBeBuiltError()
	}

	return nil, errors.NewCouldNotBeBuiltError()
}
