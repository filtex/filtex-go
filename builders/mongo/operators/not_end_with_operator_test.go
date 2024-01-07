package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNotEndWithExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeString, "Value", value)

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
	assert.Equal(t, value+"$", regex.(string))

	options, ok := notInner["$options"]
	assert.True(t, ok)
	assert.NotNil(t, options)
	assert.Equal(t, "i", options.(string))
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotEndWithExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotEndWithOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
