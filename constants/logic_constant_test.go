package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic_ToTokenType_ShouldReturnNoneToken_WhenValueIsNotValid(t *testing.T) {
	// Arrange
	samples := []Logic{
		LogicUnknown,
	}

	for _, v := range samples {
		// Act
		result := v.ToTokenType()

		// Assert
		assert.Equal(t, TokenTypeNone, result)
	}
}

func TestLogic_ToTokenType_ShouldReturnToken_WhenValueIsValid(t *testing.T) {
	// Arrange
	samples := map[Logic]TokenType{
		LogicAnd: TokenTypeAnd,
		LogicOr:  TokenTypeOr,
	}

	for k, v := range samples {
		// Act
		result := k.ToTokenType()

		// Assert
		assert.Equal(t, v, result)
	}
}

func TestLogic_ParseLogic_ShouldReturnUnknown_WhenValueIsNotValid(t *testing.T) {
	// Arrange
	samples := []string{
		"",
		"test",
	}

	for _, v := range samples {
		// Act
		result := ParseLogic(v)

		// Assert
		assert.Equal(t, LogicUnknown, result)
	}
}

func TestLogic_ParseLogic_ShouldReturnLogic_WhenValueIsValid(t *testing.T) {
	// Arrange
	samples := map[string]Logic{
		"and": LogicAnd,
		"And": LogicAnd,
		"AND": LogicAnd,
		"or":  LogicOr,
		"Or":  LogicOr,
		"OR":  LogicOr,
	}

	for k, v := range samples {
		// Act
		result := ParseLogic(k)

		// Assert
		assert.Equal(t, v, result)
	}
}
