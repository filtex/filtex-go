package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenType_ToOperator_ShouldReturnUnknown_WhenValueIsNotValid(t *testing.T) {
	// Arrange
	samples := []TokenType{
		TokenTypeNone,
		TokenTypeComma,
		TokenTypeSlash,
		TokenTypeField,
		TokenTypeValue,
		TokenTypeStringValue,
		TokenTypeNumberValue,
		TokenTypeBooleanValue,
		TokenTypeDateValue,
		TokenTypeTimeValue,
		TokenTypeDateTimeValue,
		TokenTypeLiteral,
		TokenTypeAnd,
		TokenTypeOr,
		TokenTypeSpace,
		TokenTypeOpenBracket,
		TokenTypeCloseBracket,
	}

	for _, v := range samples {
		// Act
		result := v.ToOperator()

		// Assert
		assert.Equal(t, OperatorUnknown, result)
	}
}

func TestTokenType_ToOperator_ShouldReturnOperator_WhenValueIsValid(t *testing.T) {
	// Arrange
	samples := map[TokenType]Operator{
		TokenTypeEqual:              OperatorEqual,
		TokenTypeNotEqual:           OperatorNotEqual,
		TokenTypeGreaterThan:        OperatorGreaterThan,
		TokenTypeGreaterThanOrEqual: OperatorGreaterThanOrEqual,
		TokenTypeLessThan:           OperatorLessThan,
		TokenTypeLessThanOrEqual:    OperatorLessThanOrEqual,
		TokenTypeBlank:              OperatorBlank,
		TokenTypeNotBlank:           OperatorNotBlank,
		TokenTypeContain:            OperatorContain,
		TokenTypeNotContain:         OperatorNotContain,
		TokenTypeStartWith:          OperatorStartWith,
		TokenTypeNotStartWith:       OperatorNotStartWith,
		TokenTypeEndWith:            OperatorEndWith,
		TokenTypeNotEndWith:         OperatorNotEndWith,
		TokenTypeIn:                 OperatorIn,
		TokenTypeNotIn:              OperatorNotIn,
	}

	for k, v := range samples {
		// Act
		result := k.ToOperator()

		// Assert
		assert.Equal(t, v, result)
	}
}
