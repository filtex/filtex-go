package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndNilAndValueIsEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndNilAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndEmpty(t *testing.T) {
	// Arrange
	names := make([]string, 0)
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndNil(t *testing.T) {
	// Arrange
	var names []string
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	names := []string{
		"Dog",
	}
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	names := []string{
		"Filtex",
	}
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndEmpty(t *testing.T) {
	// Arrange
	value := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 1000)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndNil(t *testing.T) {
	// Arrange
	var value []int
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 1000)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	value := []float64{
		100,
	}
	data := utils.ObjectToMap(struct {
		Value []float64
	}{
		Value: value,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 200)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	value := []float64{
		100,
	}
	data := utils.ObjectToMap(struct {
		Value []float64
	}{
		Value: value,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]bool, 0)
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndNil(t *testing.T) {
	// Arrange
	var values []bool
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	values := []bool{
		true,
	}
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	values := []bool{
		true,
	}
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	values := []time.Time{
		value,
	}
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateArray, "Values", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	value := time.Now()
	values := []time.Time{
		value,
	}
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateArray, "Values", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []int
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	values := []int{
		10,
	}
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", "00:20:00")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	values := []int{
		60,
	}
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	values := []time.Time{
		value,
	}
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	value := time.Now()
	values := []time.Time{
		value,
	}
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeNumber, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := NotContainOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
