package mongo

import (
	"testing"

	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/expressions"
	"github.com/stretchr/testify/assert"
)

func TestBuild_ShouldReturnError_WhenExpressionIsNil(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()

	// Act
	expression, err := builder.Build(nil)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestBuild_ShouldReturnError_WhenExpressionIsNotValid(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()

	// Act
	expression, err := builder.Build(struct{}{})

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestBuild_ShouldReturnError_WhenExpressionIsLogicExpressionAndNotValidLogic(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()
	logicExpression := expressions.NewLogicExpression("", []expressions.Expression{})

	// Act
	expression, err := builder.Build(logicExpression)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestBuild_ShouldReturnExpression_WhenExpressionIsLogicExpressionAndValid(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()
	logicExpression := expressions.NewLogicExpression(constants.LogicAnd, []expressions.Expression{
		expressions.NewOperatorExpression(constants.FieldTypeString, "Value", constants.OperatorEqual, "Filtex"),
	})

	// Act
	expression, err := builder.Build(logicExpression)

	// Assert
	var mongoExpression *types.MongoExpression
	assert.NotNil(t, expression)
	assert.IsType(t, mongoExpression, expression)
	assert.NoError(t, err)
}

func TestBuild_ShouldReturnError_WhenExpressionIsOperatorExpressionAndNotValidOperator(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()
	operatorExpression := expressions.NewOperatorExpression(constants.FieldTypeString, "Value", constants.Operator{}, "Filtex")

	// Act
	expression, err := builder.Build(operatorExpression)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestBuild_ShouldReturnExpression_WhenExpressionIsOperatorExpressionAndValid(t *testing.T) {
	// Arrange
	builder := NewMongoFilterBuilder()
	operatorExpression := expressions.NewOperatorExpression(constants.FieldTypeString, "Value", constants.OperatorEqual, "Filtex")

	// Act
	expression, err := builder.Build(operatorExpression)

	// Assert
	var mongoExpression *types.MongoExpression
	assert.NotNil(t, expression)
	assert.IsType(t, mongoExpression, expression)
	assert.NoError(t, err)
}
