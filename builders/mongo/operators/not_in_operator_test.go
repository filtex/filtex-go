package operators

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNotInExpression_ShouldReturnExpression_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	item := "Filtex"
	value := []interface{}{
		item,
	}

	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeString, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	nin, ok := inner["$nin"]
	assert.True(t, ok)
	assert.NotNil(t, nin)

	items, ok := nin.([]interface{})
	assert.True(t, ok)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0].(string))
}

func TestNotInExpression_ShouldReturnExpression_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	item := float64(100)
	value := []interface{}{
		item,
	}

	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	nin, ok := inner["$nin"]
	assert.True(t, ok)
	assert.NotNil(t, nin)

	items, ok := nin.([]interface{})
	assert.True(t, ok)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0].(float64))
}

func TestNotInExpression_ShouldReturnExpression_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	item := time.Now()
	value := []interface{}{
		item,
	}

	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	nin, ok := inner["$nin"]
	assert.True(t, ok)
	assert.NotNil(t, nin)

	items, ok := nin.([]interface{})
	assert.True(t, ok)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0].(time.Time))
}

func TestNotInExpression_ShouldReturnExpression_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	item := 60
	value := []interface{}{
		item,
	}

	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	nin, ok := inner["$nin"]
	assert.True(t, ok)
	assert.NotNil(t, nin)

	items, ok := nin.([]interface{})
	assert.True(t, ok)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0].(int))
}

func TestNotInExpression_ShouldReturnExpression_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	item := time.Now()
	value := []interface{}{
		item,
	}

	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Assert
	assert.NotNil(t, expression)

	field, ok := expression.Condition["Value"]
	assert.True(t, ok)
	assert.NotNil(t, field)

	inner, ok := field.(bson.M)
	assert.True(t, ok)
	assert.NotNil(t, inner)

	nin, ok := inner["$nin"]
	assert.True(t, ok)
	assert.NotNil(t, nin)

	items, ok := nin.([]interface{})
	assert.True(t, ok)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0].(time.Time))
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}

func TestNotInExpression_ShouldReturnNil_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	// Act
	expression := NotInOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Assert
	assert.Nil(t, expression)
}
