package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	regex, ok := inner["$regex"]
	assert.True(t, ok)
	assert.NotNil(t, regex)
	assert.Equal(t, value, regex.(string))

	options, ok := inner["$options"]
	assert.True(t, ok)
	assert.NotNil(t, options)
	assert.Equal(t, "i", options.(string))
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeStringArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	value := true

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeBooleanArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeDateArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeTimeArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	in, ok := inner["$in"]
	assert.True(t, ok)
	assert.NotNil(t, in)

	values, ok := in.([]interface{})
	assert.True(t, ok)
	assert.Len(t, values, 1)
	assert.Equal(t, value, values[0])
}

func TestContainExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestContainExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestContainExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestContainExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestContainExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := ContainOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
