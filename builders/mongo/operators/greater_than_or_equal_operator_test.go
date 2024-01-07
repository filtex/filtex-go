package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGreaterThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	gte, ok := inner["$gte"]
	assert.True(t, ok)
	assert.NotNil(t, gte)
	assert.Equal(t, value, gte.(float64))
}

func TestGreaterThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	gte, ok := inner["$gte"]
	assert.True(t, ok)
	assert.NotNil(t, gte)
	assert.Equal(t, value, gte.(time.Time))
}

func TestGreaterThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	gte, ok := inner["$gte"]
	assert.True(t, ok)
	assert.NotNil(t, gte)
	assert.Equal(t, value, gte.(int))
}

func TestGreaterThanOrEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	gte, ok := inner["$gte"]
	assert.True(t, ok)
	assert.NotNil(t, gte)
	assert.Equal(t, value, gte.(time.Time))
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestGreaterThanOrEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := GreaterThanOrEqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
