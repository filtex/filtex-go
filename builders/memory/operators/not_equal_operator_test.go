package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filter")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndValueIsSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueIsSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsNumberPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(101))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := float64(101)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanPointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanPointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDatePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsTimePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := NotEqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
