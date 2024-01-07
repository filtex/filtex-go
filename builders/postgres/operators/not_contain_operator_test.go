package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value NOT ILIKE '%' || $1 || '%'", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeStringArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT (LOWER($1) = ANY (LOWER(Value::TEXT)::TEXT[]))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT ($1 = ANY (Value))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	value := true

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeBooleanArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT ($1 = ANY (Value))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeDateArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT ($1 = ANY (Value))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeTimeArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT ($1 = ANY (Value))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "NOT ($1 = ANY (Value))", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestNotContainExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeNumber, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotContainExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeBoolean, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotContainExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeDate, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotContainExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestNotContainExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTime, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}
