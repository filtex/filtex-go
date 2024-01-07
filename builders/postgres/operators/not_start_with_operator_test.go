package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotStartWithExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value NOT ILIKE $1 || '%'", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeStringArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeNumber, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeBoolean, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDate, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotStartWithExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}
