package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", value)

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
	assert.Equal(t, "^"+value+"$", regex.(string))

	options, ok := inner["$options"]
	assert.True(t, ok)
	assert.NotNil(t, options)
	assert.Equal(t, "i", options.(string))
}

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, value, eq.(float64))
}

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	value := true

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, value, eq.(bool))
}

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, value, eq.(time.Time))
}

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, value, eq.(int))
}

func TestEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, value, eq.(time.Time))
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := EqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
