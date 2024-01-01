package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldType_IsArray_ShouldReturnFalse_WhenValueIsNotArray(t *testing.T) {
	// Arrange
	samples := []FieldType{
		FieldTypeUnknown,
		FieldTypeString,
		FieldTypeNumber,
		FieldTypeBoolean,
		FieldTypeDate,
		FieldTypeTime,
		FieldTypeDateTime,
	}

	for _, v := range samples {
		// Act
		result := v.IsArray()

		// Assert
		assert.False(t, result)
	}
}

func TestFieldType_IsArray_ShouldReturnTrue_WhenValueIsArray(t *testing.T) {
	// Arrange
	samples := []FieldType{
		FieldTypeStringArray,
		FieldTypeNumberArray,
		FieldTypeBooleanArray,
		FieldTypeDateArray,
		FieldTypeTimeArray,
		FieldTypeDateTimeArray,
	}

	for _, v := range samples {
		// Act
		result := v.IsArray()

		// Assert
		assert.True(t, result)
	}
}
