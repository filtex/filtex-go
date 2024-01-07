package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lte, ok := inner["$lte"]
	assert.True(t, ok)
	assert.NotNil(t, lte)
	assert.Equal(t, value, lte.(float64))
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lte, ok := inner["$lte"]
	assert.True(t, ok)
	assert.NotNil(t, lte)
	assert.Equal(t, value, lte.(time.Time))
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lte, ok := inner["$lte"]
	assert.True(t, ok)
	assert.NotNil(t, lte)
	assert.Equal(t, value, lte.(int))
}

func TestLessThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lte, ok := inner["$lte"]
	assert.True(t, ok)
	assert.NotNil(t, lte)
	assert.Equal(t, value, lte.(time.Time))
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOrEqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
