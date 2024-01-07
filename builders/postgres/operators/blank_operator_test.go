package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeString, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value IS NULL OR Value = ''", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeStringArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) = 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeNumber, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeBoolean, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDate, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}
