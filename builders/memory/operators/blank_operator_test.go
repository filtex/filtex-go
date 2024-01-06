package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})

	expression := BlankOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndEmpty(t *testing.T) {
	// Arrange
	names := make([]string, 0)
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeStringArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndNil(t *testing.T) {
	// Arrange
	var names []string
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeStringArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndEmpty(t *testing.T) {
	// Arrange
	value := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndNil(t *testing.T) {
	// Arrange
	var value []int
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]bool, 0)
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeBooleanArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndNil(t *testing.T) {
	// Arrange
	var values []bool
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeBooleanArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDateArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDateArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeTimeArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []int
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeTimeArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestBlankExpression_ShouldReturnFalse_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestBlankExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: nil,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestBlankExpression_ShouldReturnFalse_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestBlankExpression_ShouldReturnFalse_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestBlankExpression_ShouldReturnFalse_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := BlankOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
