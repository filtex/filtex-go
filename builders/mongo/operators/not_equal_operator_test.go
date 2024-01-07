package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	not, ok := inner["$not"]
	assert.True(t, ok)
	assert.NotNil(t, not)

	notInner, ok := not.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, notInner)

	regex, ok := notInner["$regex"]
	assert.True(t, ok)
	assert.NotNil(t, regex)
	assert.Equal(t, "^"+value+"$", regex.(string))

	options, ok := notInner["$options"]
	assert.True(t, ok)
	assert.NotNil(t, options)
	assert.Equal(t, "i", options.(string))
}

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	value := float64(100)

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, value, ne.(float64))
}

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	value := true

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, value, ne.(bool))
}

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, value, ne.(time.Time))
}

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	value := 60

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, value, ne.(int))
}

func TestNotEqualExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	value := time.Now()

	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	valueInner, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, valueInner)

	inner, ok := valueInner.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, value, ne.(time.Time))
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEqualExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
