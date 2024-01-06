package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := NotInOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueIsNotArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := NotInOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueContains(t *testing.T) {
	// Arrange
	value := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeString, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndValueNotContain(t *testing.T) {
	// Arrange
	value := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeString, "Value", []interface{}{"Filter"})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: float64(1000),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: float64(1000),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", float64(1000))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: float64(1000),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueContains(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndValueNotContain(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumber, "Value", []interface{}{101})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBoolean, "Value", true)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueContains(t *testing.T) {
	// Arrange
	value := false
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBoolean, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanAndValueNotContain(t *testing.T) {
	// Arrange
	value := false
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBoolean, "Value", []interface{}{true})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsNotArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueContains(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDate, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndValueNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDate, "Value", []interface{}{value.Add(24 * time.Hour)})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", 90)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueContains(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueNotContain(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTime, "Value", []interface{}{value + 1})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", now)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", now.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueContains(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTime, "Value", []interface{}{value.Add(24 * time.Hour)})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: []string{"Filtex"},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{100},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []bool
	}{
		Value: []bool{true},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{time.Now()},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{60},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestNotInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{time.Now()},
	})
	expression := NotInOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
