package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndLess(t *testing.T) {
	// Arrange
	value := float64(99)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndLess(t *testing.T) {
	// Arrange
	value := float64(99)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndEqual(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndEqual(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndGreater(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumber, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsDatePointerAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndLess(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndLess(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndEqual(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndEqual(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsTimePointerAndGreater(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndGreater(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTime, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimePointerAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: false,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []bool
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestGreaterThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := GreaterThanOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
