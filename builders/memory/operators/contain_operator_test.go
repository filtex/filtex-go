package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndNilAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringPointerAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringPointerAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value *string
	}{
		Value: &name,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndEmptyAndValueIsNotEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndEmptyAndValueIsEmpty(t *testing.T) {
	// Arrange
	name := ""
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: name,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeString, "Value", "")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringArrayAndEmpty(t *testing.T) {
	// Arrange
	names := make([]string, 0)
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringArrayAndNil(t *testing.T) {
	// Arrange
	var names []string
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsStringArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	names := []string{
		"Dog",
	}
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsStringArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	names := []string{
		"Filtex",
	}
	data := utils.ObjectToMap(struct {
		Values []string
	}{
		Values: names,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeStringArray, "Values", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArrayAndEmpty(t *testing.T) {
	// Arrange
	value := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 1000)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArrayAndNil(t *testing.T) {
	// Arrange
	var value []int
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: value,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 1000)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	value := []float64{
		100,
	}
	data := utils.ObjectToMap(struct {
		Value []float64
	}{
		Value: value,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 200)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsNumberArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	value := []float64{
		100,
	}
	data := utils.ObjectToMap(struct {
		Value []float64
	}{
		Value: value,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeNumberArray, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]bool, 0)
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArrayAndNil(t *testing.T) {
	// Arrange
	var values []bool
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	values := []bool{
		true,
	}
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	values := []bool{
		true,
	}
	data := utils.ObjectToMap(struct {
		Values []bool
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeBooleanArray, "Values", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDateArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDateArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateArrayAndNotEmptyAndNotContain(t *testing.T) {
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
	expression := ContainOperator{}.Build(constants.FieldTypeDateArray, "Values", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateArrayAndNotEmptyAndContain(t *testing.T) {
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
	expression := ContainOperator{}.Build(constants.FieldTypeDateArray, "Values", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]int, 0)
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []int
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArrayAndNotEmptyAndNotContain(t *testing.T) {
	// Arrange
	values := []int{
		10,
	}
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", "20s")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsTimeArrayAndNotEmptyAndContain(t *testing.T) {
	// Arrange
	values := []int{
		60,
	}
	data := utils.ObjectToMap(struct {
		Values []int
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeTimeArray, "Values", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArrayAndEmpty(t *testing.T) {
	// Arrange
	values := make([]time.Time, 0)
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArrayAndNil(t *testing.T) {
	// Arrange
	var values []time.Time
	data := utils.ObjectToMap(struct {
		Values []time.Time
	}{
		Values: values,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArrayAndNotEmptyAndNotContain(t *testing.T) {
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
	expression := ContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeArrayAndNotEmptyAndContain(t *testing.T) {
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
	expression := ContainOperator{}.Build(constants.FieldTypeDateTimeArray, "Values", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsNumber(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeNumber, "Value", 100)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *bool
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDate(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestContainExpression_ShouldReturnFalse_WhenFieldTypeIsDateTime(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := ContainOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
