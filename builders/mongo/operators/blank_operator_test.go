package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.True(t, exists.(bool))

	eq, ok := inner["$eq"]
	assert.True(t, ok)
	assert.NotNil(t, eq)
	assert.Equal(t, "", eq.(string))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["Value.0"]
	assert.True(t, ok)
	assert.NotNil(t, value)

	inner, ok := value.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	exists, ok := inner["$exists"]
	assert.True(t, ok)
	assert.NotNil(t, exists)
	assert.False(t, exists.(bool))
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestBlankExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := BlankOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
