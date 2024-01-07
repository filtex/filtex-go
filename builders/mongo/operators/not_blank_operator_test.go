package operators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeString, "Value", nil)

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

	ne, ok := inner["$ne"]
	assert.True(t, ok)
	assert.NotNil(t, ne)
	assert.Equal(t, "", ne.(string))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnExpression_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

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
	assert.True(t, exists.(bool))
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotBlankExpression_ShouldReturnNil_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	// Act
	expression := NotBlankOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
