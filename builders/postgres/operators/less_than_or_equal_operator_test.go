package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeNumber, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value <= $1", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDate, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value <= $1", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeTime, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value <= $1", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value, 1)

	// Assert
	assert.NotNil(t, expression)
	assert.Equal(t, "Value <= $1", expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeString, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil, 0)

	// Assert
	assert.Nil(t, expression)
}
