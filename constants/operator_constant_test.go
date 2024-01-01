package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperator_String_ShouldReturnCorrectValue(t *testing.T) {
	// Arrange
	samples := map[Operator]string{
		OperatorEqual:              "equal",
		OperatorNotEqual:           "not-equal",
		OperatorContain:            "contain",
		OperatorNotContain:         "not-contain",
		OperatorStartWith:          "start-with",
		OperatorNotStartWith:       "not-start-with",
		OperatorEndWith:            "end-with",
		OperatorNotEndWith:         "not-end-with",
		OperatorBlank:              "blank",
		OperatorNotBlank:           "not-blank",
		OperatorGreaterThan:        "greater-than",
		OperatorGreaterThanOrEqual: "greater-than-or-equal",
		OperatorLessThan:           "less-than",
		OperatorLessThanOrEqual:    "less-than-or-equal",
		OperatorIn:                 "in",
		OperatorNotIn:              "not-in",
	}

	for k, v := range samples {
		// Act
		result := k.String()

		// Assert
		assert.Equal(t, v, result)
	}
}

func TestOperator_Equals_ShouldReturnFalse_WhenValueIsNotMatched(t *testing.T) {
	// Arrange
	samples := map[Operator]string{
		OperatorEqual:      "equals",
		OperatorNotEqual:   "Not-Equals",
		OperatorContain:    "contains",
		OperatorNotContain: "Contain",
	}

	for k, v := range samples {
		// Act
		result := k.Equals(v)

		// Assert
		assert.False(t, result)
	}
}

func TestOperator_Equals_ShouldReturnTrue_WhenValueIsMatched(t *testing.T) {
	// Arrange
	samples := map[Operator]string{
		OperatorEqual:              "equal",
		OperatorNotEqual:           "Not-Equal",
		OperatorContain:            "contain",
		OperatorNotContain:         "not Contain",
		OperatorStartWith:          "start-with",
		OperatorNotStartWith:       "not-start-with",
		OperatorEndWith:            "end-with",
		OperatorNotEndWith:         "not End With",
		OperatorBlank:              "BLANK",
		OperatorNotBlank:           "not-blank",
		OperatorGreaterThan:        "greater-than",
		OperatorGreaterThanOrEqual: "greater-than-or-equal",
		OperatorLessThan:           "less-than",
		OperatorLessThanOrEqual:    "less-than-or-equal",
		OperatorIn:                 "IN",
		OperatorNotIn:              "NOT-IN",
	}

	for k, v := range samples {
		// Act
		result := k.Equals(v)

		// Assert
		assert.True(t, result)
	}
}

func TestOperator_ParseOperator_ShouldReturnUnknown_WhenValueIsNotMatched(t *testing.T) {
	// Arrange
	samples := []string{
		"Equals",
		"Contained",
		"NotEqual",
	}

	for _, v := range samples {
		// Act
		result := ParseOperator(v)

		// Assert
		assert.Equal(t, OperatorUnknown, result)
	}
}

func TestOperator_ParseOperator_ShouldReturnCorrectValue_WhenValueIsMatched(t *testing.T) {
	// Arrange
	samples := map[string]Operator{
		"equal":                 OperatorEqual,
		"Not-Equal":             OperatorNotEqual,
		"contain":               OperatorContain,
		"not Contain":           OperatorNotContain,
		"start-with":            OperatorStartWith,
		"not-start-with":        OperatorNotStartWith,
		"end-with":              OperatorEndWith,
		"not End With":          OperatorNotEndWith,
		"BLANK":                 OperatorBlank,
		"not-blank":             OperatorNotBlank,
		"greater-than":          OperatorGreaterThan,
		"greater-than-or-equal": OperatorGreaterThanOrEqual,
		"less-than":             OperatorLessThan,
		"less-than-or-equal":    OperatorLessThanOrEqual,
		"IN":                    OperatorIn,
		"NOT-IN":                OperatorNotIn,
	}

	for k, v := range samples {
		// Act
		result := ParseOperator(k)

		// Assert
		assert.Equal(t, v, result)
	}
}
