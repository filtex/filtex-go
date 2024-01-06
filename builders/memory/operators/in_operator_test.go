package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Test",
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", "Filtex")

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsStringAndValueContains(t *testing.T) {
	// Arrange
	value := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsStringAndValueNotContain(t *testing.T) {
	// Arrange
	value := "Filtex"
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", []interface{}{"Filter"})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: float64(1000),
	})
	expression := InOperator{}.Build(constants.FieldTypeNumber, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueIsNotArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: float64(1000),
	})
	expression := InOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndValueContains(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeNumber, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndValueNotContain(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeNumber, "Value", []interface{}{101})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := InOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueIsNotArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: true,
	})
	expression := InOperator{}.Build(constants.FieldTypeBoolean, "Value", false)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsBooleanAndValueContains(t *testing.T) {
	// Arrange
	value := false
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeBoolean, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanAndValueNotContain(t *testing.T) {
	// Arrange
	value := false
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeBoolean, "Value", []interface{}{true})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := InOperator{}.Build(constants.FieldTypeDate, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := InOperator{}.Build(constants.FieldTypeDate, "Value", now)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := InOperator{}.Build(constants.FieldTypeDate, "Value", now.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndValueContains(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeDate, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndValueNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeDate, "Value", []interface{}{value.Add(24 * time.Hour)})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 60,
	})
	expression := InOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: 600,
	})
	expression := InOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndValueContains(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeTime, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndValueNotContain(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeTime, "Value", []interface{}{value + 1})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: time.Now(),
	})
	expression := InOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueIsNotArrayAndContains(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := InOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueIsNotArrayAndNotContain(t *testing.T) {
	// Arrange
	now := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: now,
	})
	expression := InOperator{}.Build(constants.FieldTypeDateTime, "Value", now.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndValueContains(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeDateTime, "Value", []interface{}{value})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndValueNotContain(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := InOperator{}.Build(constants.FieldTypeDateTime, "Value", []interface{}{value.Add(24 * time.Hour)})

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: []string{"Filtex"},
	})
	expression := InOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{100},
	})
	expression := InOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []bool
	}{
		Value: []bool{true},
	})
	expression := InOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{time.Now()},
	})
	expression := InOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: []int{60},
	})
	expression := InOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestInExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: []time.Time{time.Now()},
	})
	expression := InOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
