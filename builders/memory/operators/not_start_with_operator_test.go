package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotStartWithExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndNilAndValueIsEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotStartWithExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndNilAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotStartWithExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndDoesNotEndWithValue(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "ex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEndsWithValue(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeString, "Value", "Fil")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 100,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeNumber, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{
			100,
		},
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeNumberArray, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []bool
	}{
		Value: []bool{
			true,
		},
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeBooleanArray, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{
			time.Now(),
		},
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateArray, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{
			60,
		},
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeTimeArray, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotStartWithExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{
			time.Now(),
		},
	})
	expression := NotStartWithOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
