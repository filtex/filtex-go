package operators

import (
	"github.com/filtex/filtex-go/builders/memory/utils"
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", float64(100))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsNumberPointerAndLess(t *testing.T) {
	// Arrange
	value := float64(99)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsNumberAndLess(t *testing.T) {
	// Arrange
	value := float64(99)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberPointerAndEqual(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value *float64
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndEqual(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberAndGreater(t *testing.T) {
	// Arrange
	value := float64(100)
	data := utils.ObjectToMap(struct {
		Value float64
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumber, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsDatePointerAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDatePointerAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDate, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", 60)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsTimePointerAndLess(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsTimeAndLess(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value+1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndEqual(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndEqual(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimePointerAndGreater(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value *int
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeAndGreater(t *testing.T) {
	// Arrange
	value := 60
	data := utils.ObjectToMap(struct {
		Value int
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTime, "Value", value-1)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndNil(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", time.Now())

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimePointerAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnTrue_WhenFieldTypeIsDateTimeAndLess(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.True(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndEqual(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimePointerAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value *time.Time
	}{
		Value: &value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeAndGreater(t *testing.T) {
	// Arrange
	value := time.Now()
	data := utils.ObjectToMap(struct {
		Value time.Time
	}{
		Value: value,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTime, "Value", value.Add(-24*time.Hour))

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsString(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value string
	}{
		Value: "Filtex",
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeString, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsStringArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []string
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeStringArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsBoolean(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value bool
	}{
		Value: false,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeBoolean, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsBooleanArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []bool
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeBooleanArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsNumberArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeNumberArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []int
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}

func TestLessThanExpression_ShouldReturnFalse_WhenFieldTypeIsDateTimeArray(t *testing.T) {
	// Arrange
	data := utils.ObjectToMap(struct {
		Value []time.Time
	}{
		Value: nil,
	})
	expression := LessThanOperator{}.Build(constants.FieldTypeDateTimeArray, "Value", nil)

	// Act
	result := expression.Fn(data)

	// Assert
	assert.False(t, result)
}
