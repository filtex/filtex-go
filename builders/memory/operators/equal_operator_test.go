package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filter")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndValueIsSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndValueIsSame(t *testing.T) {
	// Arrange
	name := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(101))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsNumberPointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := float64(101)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndValueIsSame(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanPointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanPointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanPointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanAndValueIsSame(t *testing.T) {
	// Arrange
	value := true
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDatePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsTimePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueIsSame(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndNilAndValueIsNotNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimePointerAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsNotSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueIsSame(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestEqualExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := EqualOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
