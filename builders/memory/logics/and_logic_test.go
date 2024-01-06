package logics

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/filtex/filtex-go/builders/memory/types"
)

func TestAndExpression_ShouldReturnTrue_WhenAllExpressionsReturnTrue(t *testing.T) {
	// Arrange
	firstExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return true
		},
	}
	secondExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return true
		},
	}
	thirdExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return true
		},
	}
	expression := AndLogic{}.Build([]*types.MemoryExpression{
		firstExpression,
		secondExpression,
		thirdExpression,
	})

	// Act
	result := expression.Fn(nil)

	// Assert
	assert.True(t, result)
}

func TestAndExpression_ShouldReturnFalse_WhenOneExpressionReturnFalse(t *testing.T) {
	// Arrange
	firstExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return true
		},
	}
	secondExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return false
		},
	}
	thirdExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return true
		},
	}
	expression := AndLogic{}.Build([]*types.MemoryExpression{
		firstExpression,
		secondExpression,
		thirdExpression,
	})

	// Act
	result := expression.Fn(nil)

	// Assert
	assert.False(t, result)
}

func TestAndExpression_ShouldReturnFalse_WhenAllExpressionsReturnFalse(t *testing.T) {
	// Arrange
	firstExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return false
		},
	}
	secondExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return false
		},
	}
	thirdExpression := &types.MemoryExpression{
		Fn: func(map[string]interface{}) bool {
			return false
		},
	}
	expression := AndLogic{}.Build([]*types.MemoryExpression{
		firstExpression,
		secondExpression,
		thirdExpression,
	})

	// Act
	result := expression.Fn(nil)

	// Assert
	assert.False(t, result)
}
