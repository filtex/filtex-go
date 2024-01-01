package validators

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTextQueryValidator_Validate_ShouldReturnNil_WhenThereIsNoToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"": {},
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

	for query, tokens := range queryMap {
		// Arrange
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryValidator := TextQueryValidator{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		err := textQueryValidator.Validate(query)

		// Assert
		assert.NoError(t, err)
	}
}

func TestTextQueryValidator_Validate_ShouldReturnError_WhenThereIsNoneToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value": {
			{
				Type:  constants.TokenTypeNone,
				Value: "Value",
			},
		},
		"Value Equals": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeNone,
				Value: "Equals",
			},
		},
		"Value Equal 123": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			{
				Type:  constants.TokenTypeNone,
				Value: "123",
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

	for query, tokens := range queryMap {
		// Arrange
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryValidator := TextQueryValidator{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		err := textQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestTextQueryValidator_Validate_ShouldReturnError_WhenBracketCountIsMismatched(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"(Value": {
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		"(Value))": {
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
			},
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		"Value))": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
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

	for query, tokens := range queryMap {
		// Arrange
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryValidator := TextQueryValidator{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		err := textQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestTextQueryValidator_Validate_ShouldReturnError_WhenLastTokenIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		"Value Equal": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
		},
		"Value Equal Filtex And": {
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
				Value: "Filtex",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
		},
		"Value Equal Filtex And (": {
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
				Value: "Filtex",
			},
			{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
			{
				Type:  constants.TokenTypeOpenBracket,
				Value: "(",
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

	for query, tokens := range queryMap {
		// Arrange
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryValidator := TextQueryValidator{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		err := textQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestTextQueryValidator_Validate_ShouldReturnNil_WhenLastTokenIsValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value Equal Filtex": {
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
				Value: "Filtex",
			},
		},
		"(Value Equal Filtex)": {
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
				Value: "Filtex",
			},
			{
				Type:  constants.TokenTypeCloseBracket,
				Value: ")",
			},
		},
		"Value Blank": {
			{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			{
				Type:  constants.TokenTypeBlank,
				Value: "Blank",
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

	for query, tokens := range queryMap {
		// Arrange
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryValidator := TextQueryValidator{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		err := textQueryValidator.Validate(query)

		// Assert
		assert.NoError(t, err)
	}
}
