package tokenizers

import (
	"testing"
	"time"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/assert"
)

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNil_WhenLastAndNewTokenIsSpace(t *testing.T) {
	// Arrange
	tokens := []models.Token{
		{
			Type:  constants.TokenTypeSpace,
			Value: " ",
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeSpace, " ")

	// Assert
	assert.Nil(t, token)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnSpaceToken_WhenNewTokenIsSpaceAndLastTokenIsNot(t *testing.T) {
	// Arrange
	tokens := []models.Token{
		{
			Type:  constants.TokenTypeField,
			Value: "Value",
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeSpace, " ")

	// Assert
	assert.NotNil(t, token)
	assert.Equal(t, constants.TokenTypeSpace, token.Type)
	assert.Equal(t, " ", token.Value)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnFieldToken_WhenThereIsNoTokensAndTokenIsValidLiteral(t *testing.T) {
	// Arrange
	var tokens []models.Token

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value")

	// Assert
	assert.NotNil(t, token)
	assert.Equal(t, constants.TokenTypeField, token.Type)
	assert.Equal(t, "Value", token.Value)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnFieldToken_WhenThereIsNoTokensAndTokenIsValidField(t *testing.T) {
	// Arrange
	var tokens []models.Token

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value")

	// Assert
	assert.NotNil(t, token)
	assert.Equal(t, constants.TokenTypeField, token.Type)
	assert.Equal(t, "Value", token.Value)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenThereIsNoTokensAndTokenIsInvalidField(t *testing.T) {
	// Arrange
	var tokens []models.Token

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value1")

	// Assert
	assert.NotNil(t, token)
	assert.Equal(t, constants.TokenTypeNone, token.Type)
	assert.Equal(t, "Value1", token.Value)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnFieldToken_WhenThereIsNoTokensAndTokenIsInOpenGroupTokens(t *testing.T) {
	// Arrange
	var tokens []models.Token

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeOpenBracket, "(")

	// Assert
	assert.NotNil(t, token)
	assert.Equal(t, constants.TokenTypeOpenBracket, token.Type)
	assert.Equal(t, "(", token.Value)
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsFieldAndLastTokenIsNotInPreFieldAndComparerAndSeparatorTokensAndValidForValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeValue, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsFieldAndLastTokenIsNotInPreFieldAndComparerAndSeparatorTokensAndNotValidForValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeNone,
				Value: "",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: false,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsFieldAndLastTokenIsInPreFieldTokensAndNotValidField(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value1")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value1", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnFieldToken_WhenTokenIsFieldAndLastTokenIsInPreFieldTokensAndValidField(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeField, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsFieldAndLastTokenIsInComparerOrSeparatorTokensAndInvalidValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeNumber.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Test")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Test", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenLastTokenIsInComparerOrSeparatorTokensAndValidValueAndOperatorIsNotComparer(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsFieldAndLastTokenIsInComparerOrSeparatorTokensAndValidValueAndOperatorIsComparer(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeField, "Test")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeValue, token.Type)
		assert.Equal(t, "Test", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsLiteralAndLastTokenIsNotInComparerAndSeparatorAndPreFieldTokensAndValidForValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeValue, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLiteralAndLastTokenIsNotInComparerAndSeparatorAndPreFieldTokensAndNotValidForValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeNone,
				Value: "",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: false,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLiteralAndLastTokenIsInComparerOrSeparatorTokensAndInvalidValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeDate.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "100")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "100", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLiteralAndLastTokenIsInComparerOrSeparatorTokensAndValidValueAndLastOperatorIsNotComparer(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Test")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Test", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsLiteralAndLastTokenIsInComparerOrSeparatorTokensAndValidValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Test")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeValue, token.Type)
		assert.Equal(t, "Test", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLiteralAndLastTokenIsNotInComparerAndSeparatorAndPreFieldTokensAndInvalidForLookupValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeBoolean.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: []models.Lookup{
					{
						Name:  "Enabled",
						Value: true,
					},
					{
						Name:  "Disabled",
						Value: false,
					},
				},
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsLiteralAndLastTokenIsNotInComparerAndSeparatorAndPreFieldTokensAndValidForLookupValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeBoolean.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: []models.Lookup{
					{
						Name:  "Enabled",
						Value: true,
					},
					{
						Name:  "Disabled",
						Value: false,
					},
				},
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Enabled")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeValue, token.Type)
		assert.Equal(t, true, token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLiteralAndLastTokenIsInPreFieldTokensAndInvalidField(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value1")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "Value1", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnFieldToken_WhenTokenIsLiteralAndLastTokenIsInPreFieldTokensAndValidField(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeLiteral, "Value")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeField, token.Type)
		assert.Equal(t, "Value", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsValueAndLastTokenIsNotInComparerAndSeparatorTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeNone,
				Value: "",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: false,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
			},
		},
	}

	valueTokenTypes := []constants.TokenType{
		constants.TokenTypeValue,
		constants.TokenTypeStringValue,
		constants.TokenTypeNumberValue,
		constants.TokenTypeBooleanValue,
		constants.TokenTypeDateValue,
		constants.TokenTypeTimeValue,
		constants.TokenTypeDateTimeValue,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, valueTokenType := range valueTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, valueTokenType, "Test")

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, "Test", token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsValueAndLastTokenIsInComparerOrSeparatorTokensAndInvalidValue(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	valueTokenTypes := []constants.TokenType{
		constants.TokenTypeValue,
		constants.TokenTypeStringValue,
		constants.TokenTypeNumberValue,
		constants.TokenTypeBooleanValue,
		constants.TokenTypeDateValue,
		constants.TokenTypeTimeValue,
		constants.TokenTypeDateTimeValue,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeTime.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, valueTokenType := range valueTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, valueTokenType, "TEST")

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, "TEST", token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsValueAndLastTokenIsInComparerOrSeparatorTokensAndValidValueAndOperatorIsNotComparer(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	valueTokenTypes := []constants.TokenType{
		constants.TokenTypeValue,
		constants.TokenTypeStringValue,
		constants.TokenTypeNumberValue,
		constants.TokenTypeBooleanValue,
		constants.TokenTypeDateValue,
		constants.TokenTypeTimeValue,
		constants.TokenTypeDateTimeValue,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, valueTokenType := range valueTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, valueTokenType, "100")

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, "100", token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnValueToken_WhenTokenIsValueAndLastTokenIsInComparerOrSeparatorTokensAndValidValueAndOperatorIsComparer(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	valueTokenTypes := []constants.TokenType{
		constants.TokenTypeValue,
		constants.TokenTypeStringValue,
		constants.TokenTypeNumberValue,
		constants.TokenTypeBooleanValue,
		constants.TokenTypeDateValue,
		constants.TokenTypeTimeValue,
		constants.TokenTypeDateTimeValue,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, valueTokenType := range valueTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, valueTokenType, "Test")

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, valueTokenType, token.Type)
			assert.Equal(t, "Test", token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsOperatorAndLastTokenIsNotFieldToken(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	operatorTokenTypes := []constants.TokenType{
		constants.TokenTypeEqual,
		constants.TokenTypeNotEqual,
		constants.TokenTypeGreaterThan,
		constants.TokenTypeGreaterThanOrEqual,
		constants.TokenTypeLessThan,
		constants.TokenTypeLessThanOrEqual,
		constants.TokenTypeBlank,
		constants.TokenTypeNotBlank,
		constants.TokenTypeContain,
		constants.TokenTypeNotContain,
		constants.TokenTypeStartWith,
		constants.TokenTypeNotStartWith,
		constants.TokenTypeEndWith,
		constants.TokenTypeNotEndWith,
		constants.TokenTypeIn,
		constants.TokenTypeNotIn,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, operatorTokenType := range operatorTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, operatorTokenType, string(operatorTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, string(operatorTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsOperatorAndLastTokenIsFieldTokenAndInvalidOperator(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
	}

	invalidOperators := []string{
		"",
		"op",
		string(constants.TokenTypeGreaterThan),
		string(constants.TokenTypeGreaterThanOrEqual),
		string(constants.TokenTypeLessThan),
		string(constants.TokenTypeLessThanOrEqual),
		string(constants.TokenTypeBlank),
		string(constants.TokenTypeNotBlank),
		string(constants.TokenTypeContain),
		string(constants.TokenTypeNotContain),
		string(constants.TokenTypeStartWith),
		string(constants.TokenTypeNotStartWith),
		string(constants.TokenTypeEndWith),
		string(constants.TokenTypeNotEndWith),
		string(constants.TokenTypeIn),
		string(constants.TokenTypeNotIn),
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, invalidOperator := range invalidOperators {
			// Act
			token := baseQueryTokenizer.createToken(tokens, constants.TokenType(invalidOperator), invalidOperator)

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, invalidOperator, token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnOperatorToken_WhenTokenIsOperatorAndLastTokenIsFieldTokenAndValidOperator(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
	}

	operatorTokenTypes := []constants.TokenType{
		constants.TokenTypeEqual,
		constants.TokenTypeNotEqual,
		constants.TokenTypeGreaterThan,
		constants.TokenTypeGreaterThanOrEqual,
		constants.TokenTypeLessThan,
		constants.TokenTypeLessThanOrEqual,
		constants.TokenTypeBlank,
		constants.TokenTypeNotBlank,
		constants.TokenTypeContain,
		constants.TokenTypeNotContain,
		constants.TokenTypeStartWith,
		constants.TokenTypeNotStartWith,
		constants.TokenTypeEndWith,
		constants.TokenTypeNotEndWith,
		constants.TokenTypeIn,
		constants.TokenTypeNotIn,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, operatorTokenType := range operatorTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, operatorTokenType, string(operatorTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, operatorTokenType, token.Type)
			assert.Equal(t, string(operatorTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsLogicAndLastTokenIsNotInValueAndCloseGroupAndNotComparerTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
		},
	}

	logicTokenTypes := []constants.TokenType{
		constants.TokenTypeAnd,
		constants.TokenTypeOr,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, logicTokenType := range logicTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, logicTokenType, string(logicTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, string(logicTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnLogicToken_WhenTokenIsLogicAndLastTokenIsInValueOrCloseGroupOrNotComparerTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test1",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: true,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
	}

	logicTokenTypes := []constants.TokenType{
		constants.TokenTypeAnd,
		constants.TokenTypeOr,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, logicTokenType := range logicTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, logicTokenType, string(logicTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, logicTokenType, token.Type)
			assert.Equal(t, string(logicTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsInOpenGroupTokenTypesAndLastTokenIsNotInLogicAndOpenGroupTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeOpenBracket, "(")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, "(", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnOpenGroupToken_WhenTokenIsInOpenGroupTokensAndLastTokenIsInLogicOrOpenGroupTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
			{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeOpenBracket, "(")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeOpenBracket, token.Type)
		assert.Equal(t, "(", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsInCloseGroupTokenTypesAndLastTokenIsNotInValueAndCloseGroupAndNotComparerTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeSlash,
				Value: "/",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeCloseBracket, ")")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, ")", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsInOpenGroupTokensAndLastTokenIsInValueOrCloseGroupAndNotComparerTokensAndOpenGroupCountIsInvalid(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: false,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeCloseBracket, ")")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeNone, token.Type)
		assert.Equal(t, ")", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnCloseGroupToken_WhenTokenIsInOpenGroupTokensAndLastTokenIsInValueOrCloseGroupAndNotComparerTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeBooleanValue,
				Value: false,
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateValue,
				Value: time.Now(),
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeTimeValue,
				Value: 60,
			},
		},
		{
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeDateTimeValue,
				Value: time.Now(),
			},
		},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		// Act
		token := baseQueryTokenizer.createToken(tokens, constants.TokenTypeCloseBracket, ")")

		// Assert
		assert.NotNil(t, token)
		assert.Equal(t, constants.TokenTypeCloseBracket, token.Type)
		assert.Equal(t, ")", token.Value)
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsSeparatorAndOperatorIsNotInComparerAndMultiAllowedTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEqual,
				Value: "Not Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThan,
				Value: "Greater Than",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeGreaterThanOrEqual,
				Value: "Greater Than Or Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThan,
				Value: "Less Than",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeLessThanOrEqual,
				Value: "Less Than Or Equal",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotBlank,
				Value: "Not Blank",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeContain,
				Value: "Contain",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotContain,
				Value: "Not Contain",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeStartWith,
				Value: "Start With",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotStartWith,
				Value: "Not Start With",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEndWith,
				Value: "End With",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotEndWith,
				Value: "Not End With",
			},
			{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
	}

	separatorTokenTypes := []constants.TokenType{
		constants.TokenTypeComma,
		constants.TokenTypeSlash,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, separatorTokenType := range separatorTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, separatorTokenType, string(separatorTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, string(separatorTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnNoneToken_WhenTokenIsSeparatorAndOperatorIsValidAndLastTokenIsNotInValueTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
			{
				Type:  constants.TokenTypeNone,
				Value: 100,
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
			{
				Type:  constants.TokenTypeComma,
				Value: ",",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
			{
				Type:  constants.TokenTypeNumberValue,
				Value: 100,
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
	}

	separatorTokenTypes := []constants.TokenType{
		constants.TokenTypeComma,
		constants.TokenTypeSlash,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, separatorTokenType := range separatorTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, separatorTokenType, string(separatorTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, constants.TokenTypeNone, token.Type)
			assert.Equal(t, string(separatorTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_CreateToken_ShouldReturnSeparatorToken_WhenTokenIsSeparatorAndOperatorIsValidAndLastTokenIsInValueTokens(t *testing.T) {
	// Arrange
	tokensList := [][]models.Token{
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeIn,
				Value: "In",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		{
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNotIn,
				Value: "Not In",
			},
			{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
	}

	separatorTokenTypes := []constants.TokenType{
		constants.TokenTypeComma,
		constants.TokenTypeSlash,
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
					constants.OperatorGreaterThan.String(),
					constants.OperatorGreaterThanOrEqual.String(),
					constants.OperatorLessThan.String(),
					constants.OperatorLessThanOrEqual.String(),
					constants.OperatorBlank.String(),
					constants.OperatorNotBlank.String(),
					constants.OperatorContain.String(),
					constants.OperatorNotContain.String(),
					constants.OperatorStartWith.String(),
					constants.OperatorNotStartWith.String(),
					constants.OperatorEndWith.String(),
					constants.OperatorNotEndWith.String(),
					constants.OperatorIn.String(),
					constants.OperatorNotIn.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for _, tokens := range tokensList {
		for _, separatorTokenType := range separatorTokenTypes {
			// Act
			token := baseQueryTokenizer.createToken(tokens, separatorTokenType, string(separatorTokenType))

			// Assert
			assert.NotNil(t, token)
			assert.Equal(t, separatorTokenType, token.Type)
			assert.Equal(t, string(separatorTokenType), token.Value)
		}
	}
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnNil_WhenNotMatched(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	// Act
	match := baseQueryTokenizer.findMatch("%")

	// Assert
	assert.Nil(t, match)
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnTokenMatch_WhenMatchedExactly(t *testing.T) {
	// Arrange
	matchMap := map[string]tokenMatch{
		"(":                     {"", constants.TokenTypeOpenBracket, "("},
		")":                     {"", constants.TokenTypeCloseBracket, ")"},
		",":                     {"", constants.TokenTypeComma, ","},
		"/":                     {"", constants.TokenTypeSlash, "/"},
		"and":                   {"", constants.TokenTypeAnd, "and"},
		"&&":                    {"", constants.TokenTypeAnd, "&&"},
		"or":                    {"", constants.TokenTypeOr, "or"},
		"||":                    {"", constants.TokenTypeOr, "||"},
		"=":                     {"", constants.TokenTypeEqual, "="},
		"equal":                 {"", constants.TokenTypeEqual, "equal"},
		"!=":                    {"", constants.TokenTypeNotEqual, "!="},
		"not equal":             {"", constants.TokenTypeNotEqual, "not equal"},
		"greater than or equal": {"", constants.TokenTypeGreaterThanOrEqual, "greater than or equal"},
		">=":                    {"", constants.TokenTypeGreaterThanOrEqual, ">="},
		"greater than":          {"", constants.TokenTypeGreaterThan, "greater than"},
		">":                     {"", constants.TokenTypeGreaterThan, ">"},
		"less than or equal":    {"", constants.TokenTypeLessThanOrEqual, "less than or equal"},
		"<=":                    {"", constants.TokenTypeLessThanOrEqual, "<="},
		"less than":             {"", constants.TokenTypeLessThan, "less than"},
		"<":                     {"", constants.TokenTypeLessThan, "<"},
		"[]":                    {"", constants.TokenTypeBlank, "[]"},
		"blank":                 {"", constants.TokenTypeBlank, "blank"},
		"![]":                   {"", constants.TokenTypeNotBlank, "![]"},
		"not blank":             {"", constants.TokenTypeNotBlank, "not blank"},
		"~":                     {"", constants.TokenTypeContain, "~"},
		"contain":               {"", constants.TokenTypeContain, "contain"},
		"!~":                    {"", constants.TokenTypeNotContain, "!~"},
		"not contain":           {"", constants.TokenTypeNotContain, "not contain"},
		"~*":                    {"", constants.TokenTypeStartWith, "~*"},
		"start with":            {"", constants.TokenTypeStartWith, "start with"},
		"!~*":                   {"", constants.TokenTypeNotStartWith, "!~*"},
		"not start with":        {"", constants.TokenTypeNotStartWith, "not start with"},
		"*~":                    {"", constants.TokenTypeEndWith, "*~"},
		"end with":              {"", constants.TokenTypeEndWith, "end with"},
		"!*~":                   {"", constants.TokenTypeNotEndWith, "!*~"},
		"not end with":          {"", constants.TokenTypeNotEndWith, "not end with"},
		"in":                    {"", constants.TokenTypeIn, "in"},
		"not in":                {"", constants.TokenTypeNotIn, "not in"},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for query, expected := range matchMap {
		// Act
		match := baseQueryTokenizer.findMatch(query)

		// Assert
		assert.Equal(t, expected.tokenType, match.tokenType)
		assert.Equal(t, expected.value, match.value)
		assert.Equal(t, expected.remainingText, match.remainingText)
	}
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnTokenMatch_WhenMatchedWithCaseInsensitively(t *testing.T) {
	// Arrange
	matchMap := map[string]tokenMatch{
		"AnD":                   {"", constants.TokenTypeAnd, "AnD"},
		"OR":                    {"", constants.TokenTypeOr, "OR"},
		"EquAL":                 {"", constants.TokenTypeEqual, "EquAL"},
		"not EQUAL":             {"", constants.TokenTypeNotEqual, "not EQUAL"},
		"greater than or EQUAL": {"", constants.TokenTypeGreaterThanOrEqual, "greater than or EQUAL"},
		"GREATER than":          {"", constants.TokenTypeGreaterThan, "GREATER than"},
		"less THAN OR equal":    {"", constants.TokenTypeLessThanOrEqual, "less THAN OR equal"},
		"less Than":             {"", constants.TokenTypeLessThan, "less Than"},
		"BlAnk":                 {"", constants.TokenTypeBlank, "BlAnk"},
		"not BlAnk":             {"", constants.TokenTypeNotBlank, "not BlAnk"},
		"Contain":               {"", constants.TokenTypeContain, "Contain"},
		"not Contain":           {"", constants.TokenTypeNotContain, "not Contain"},
		"Start with":            {"", constants.TokenTypeStartWith, "Start with"},
		"Not Start With":        {"", constants.TokenTypeNotStartWith, "Not Start With"},
		"end WiTh":              {"", constants.TokenTypeEndWith, "end WiTh"},
		"not END with":          {"", constants.TokenTypeNotEndWith, "not END with"},
		"IN":                    {"", constants.TokenTypeIn, "IN"},
		"nOt iN":                {"", constants.TokenTypeNotIn, "nOt iN"},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for query, expected := range matchMap {
		// Act
		match := baseQueryTokenizer.findMatch(query)

		// Assert
		assert.Equal(t, expected.tokenType, match.tokenType)
		assert.Equal(t, expected.value, match.value)
		assert.Equal(t, expected.remainingText, match.remainingText)
	}
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnTokenMatch_WhenMatchedAndThereIsRemainingText(t *testing.T) {
	// Arrange
	matchMap := map[string]tokenMatch{
		"Equal Test":     {" Test", constants.TokenTypeEqual, "Equal"},
		"Equal Test AND": {" Test AND", constants.TokenTypeEqual, "Equal"},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for query, expected := range matchMap {
		// Act
		match := baseQueryTokenizer.findMatch(query)

		// Assert
		assert.Equal(t, expected.tokenType, match.tokenType)
		assert.Equal(t, expected.value, match.value)
		assert.Equal(t, expected.remainingText, match.remainingText)
	}
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnTokenMatch_WhenMatchedDefinedFields(t *testing.T) {
	// Arrange
	matchMap := map[string]tokenMatch{
		"Value": {"", constants.TokenTypeField, "Value"},
		"VaLue": {"", constants.TokenTypeField, "VaLue"},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for query, expected := range matchMap {
		// Act
		match := baseQueryTokenizer.findMatch(query)

		// Assert
		assert.Equal(t, expected.tokenType, match.tokenType)
		assert.Equal(t, expected.value, match.value)
		assert.Equal(t, expected.remainingText, match.remainingText)
	}
}

func TestBaseQueryTokenizer_FindMatch_ShouldReturnTokenMatch_WhenMatchedValue(t *testing.T) {
	// Arrange
	matchMap := map[string]tokenMatch{
		"Test":                {"", constants.TokenTypeLiteral, "Test"},
		"'Test'":              {"", constants.TokenTypeStringValue, "'Test'"},
		"123":                 {"", constants.TokenTypeNumberValue, "123"},
		"True":                {"", constants.TokenTypeBooleanValue, "True"},
		"2020-01-01":          {"", constants.TokenTypeDateValue, "2020-01-01"},
		"14:12:10":            {"", constants.TokenTypeTimeValue, "14:12:10"},
		"2020-01-01 00:00:00": {"", constants.TokenTypeDateTimeValue, "2020-01-01 00:00:00"},
	}

	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	for query, expected := range matchMap {
		// Act
		match := baseQueryTokenizer.findMatch(query)

		// Assert
		assert.Equal(t, expected.tokenType, match.tokenType)
		assert.Equal(t, expected.value, match.value)
		assert.Equal(t, expected.remainingText, match.remainingText)
	}
}

func TestBaseQueryTokenizer_ValidateField_ShouldReturnFalse_WhenThereIsNoDefinedField(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"",
		"Value",
		"Value1",
		"Value2",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateField(sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateField_ShouldReturnFalse_WhenFieldIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"",
		"Value1",
		"Value2",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateField(sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateField_ShouldReturnTrue_WhenFieldIsValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Value",
		"value",
		"VaLuE",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateField(sample)

		// Assert
		assert.True(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateOperator_ShouldReturnFalse_WhenThereIsNoDefinedField(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"equal",
		"Equal",
		"not-equal",
		"Not-Equal",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateOperator("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateOperator_ShouldReturnFalse_WhenFieldIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"equal",
		"Equal",
		"not-equal",
		"Not-Equal",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateOperator("Value1", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateOperator_ShouldReturnFalse_WhenFieldIsValidAndThereIsNoDefinedOperator(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:      "Value",
				Type:      "string",
				Label:     "Value",
				Operators: []string{},
				Values:    nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"equal",
		"Equal",
		"not-equal",
		"Not-Equal",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateOperator("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateOperator_ShouldReturnFalse_WhenFieldIsValidAndOperatorIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"In",
		"Not-In",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateOperator("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateOperator_ShouldReturnTrue_WhenFieldIsValidAndOperatorIsValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"equal",
		"Equal",
		"not-equal",
		"Not-Equal",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateOperator("Value", sample)

		// Assert
		assert.True(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnFalse_WhenThereIsNoDefinedFields(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Test1",
		"Test2",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnFalse_WhenFieldIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Test1",
		"Test2",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value1", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnFalse_WhenFieldIsValidAndHasValuesAndNotMatched(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: []models.Lookup{
					{
						Name:  "Enabled",
						Value: true,
					},
					{
						Name:  "Disabled",
						Value: false,
					},
				},
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Active",
		"Passive",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnTrue_WhenFieldIsValidAndHasValuesAndMatched(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: []models.Lookup{
					{
						Name:  "Enabled",
						Value: true,
					},
					{
						Name:  "Disabled",
						Value: false,
					},
				},
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Enabled",
		"Disabled",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value", sample)

		// Assert
		assert.True(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnFalse_WhenFieldIsValidAndDoesNotHaveValuesAndTypeIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeTime.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"True",
		"2020-01-01 11:12:13",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value", sample)

		// Assert
		assert.False(t, result)
	}
}

func TestBaseQueryTokenizer_ValidateValue_ShouldReturnTrue_WhenFieldIsValidAndDoesNotHaveValuesAndTypeIsValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"Test1",
		"Test2",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.validateValue("Value", sample)

		// Assert
		assert.True(t, result)
	}
}

func TestBaseQueryTokenizer_CastValue_ShouldReturnSame_WhenThereIsNoDefinedField(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"100",
		"200",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.castValue("Value", sample)

		// Assert
		assert.Equal(t, sample, result)
	}
}

func TestBaseQueryTokenizer_CastValue_ShouldReturnSame_WhenFieldIsNotValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "Value",
				Type:  constants.FieldTypeString.String(),
				Label: "Value",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	samples := []string{
		"100",
		"200",
	}

	for _, sample := range samples {
		// Act
		result := baseQueryTokenizer.castValue("Value1", sample)

		// Assert
		assert.Equal(t, sample, result)
	}
}

func TestBaseQueryTokenizer_CastValue_ShouldReturnCasted_WhenFieldIsValid(t *testing.T) {
	// Arrange
	metadata := models.Metadata{
		Fields: []models.Field{
			{
				Name:  "StringValue",
				Type:  constants.FieldTypeString.String(),
				Label: "StringValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "StringArrayValue",
				Type:  "string-array",
				Label: "StringArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "NumberValue",
				Type:  "number",
				Label: "NumberValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "NumberArrayValue",
				Type:  "number-array",
				Label: "NumberArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "BooleanValue",
				Type:  "boolean",
				Label: "BooleanValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "BooleanArrayValue",
				Type:  "boolean-array",
				Label: "BooleanArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "DateValue",
				Type:  "date",
				Label: "DateValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "DateArrayValue",
				Type:  "date-array",
				Label: "DateArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "TimeValue",
				Type:  "time",
				Label: "TimeValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "TimeArrayValue",
				Type:  "time-array",
				Label: "TimeArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "DateTimeValue",
				Type:  "datetime",
				Label: "DateTimeValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "DateTimeArrayValue",
				Type:  "datetime-array",
				Label: "DateTimeArrayValue",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
		},
	}

	baseQueryTokenizer := NewBaseQueryTokenizer(&metadata)

	dateValueResult := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	timeValueResult := 1*60*60 + 15*60
	dateTimeValueResult := time.Date(2020, 01, 01, 11, 12, 13, 0, time.UTC)

	samples := map[string][2]interface{}{
		"StringValue":        {100, "100"},
		"StringArrayValue":   {100, "100"},
		"NumberValue":        {"100", float64(100)},
		"NumberArrayValue":   {"100", float64(100)},
		"BooleanValue":       {"True", true},
		"BooleanArrayValue":  {"False", false},
		"DateValue":          {"2020-01-01", &dateValueResult},
		"DateArrayValue":     {"2020-01-01", &dateValueResult},
		"TimeValue":          {"1h15m", &timeValueResult},
		"TimeArrayValue":     {"1h15m", &timeValueResult},
		"DateTimeValue":      {"2020-01-01 11:12:13", &dateTimeValueResult},
		"DateTimeArrayValue": {"2020-01-01 11:12:13", &dateTimeValueResult},
	}

	for field, sample := range samples {
		// Act
		result := baseQueryTokenizer.castValue(field, sample[0])

		// Assert
		assert.Equal(t, sample[1], result)
	}
}
