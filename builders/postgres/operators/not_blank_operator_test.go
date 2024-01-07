package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeString, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value IS NOT NULL AND Value <> ''", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeStringArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil, 0)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "ARRAY_LENGTH(Value, 1) <> 0", expression.Condition)
	assert.Empty(t, expression.Args)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeNumber, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeBoolean, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDate, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}
