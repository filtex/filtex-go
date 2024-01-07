package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLessThanExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lt, ok := inner["$lt"]
	assert.True(t, ok)
	assert.NotNil(t, lt)
	assert.Equal(t, value, lt.(float64))
}

func TestLessThanExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lt, ok := inner["$lt"]
	assert.True(t, ok)
	assert.NotNil(t, lt)
	assert.Equal(t, value, lt.(time.Time))
}

func TestLessThanExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lt, ok := inner["$lt"]
	assert.True(t, ok)
	assert.NotNil(t, lt)
	assert.Equal(t, value, lt.(int))
}

func TestLessThanExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	lt, ok := inner["$lt"]
	assert.True(t, ok)
	assert.NotNil(t, lt)
	assert.Equal(t, value, lt.(time.Time))
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestLessThanExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
