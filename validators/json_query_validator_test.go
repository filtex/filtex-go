package validators

import (
	"errors"
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJsonQueryValidator_Validate_ShouldReturnError_WhenQueryTokenizerReturnedError(t *testing.T) {
	// Arrange
	query := "some_text"

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
	jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()
	jsonQueryTokenizerMock.
		On("Tokenize", mock.Anything).
		Return(nil, errors.New("some error"))

	jsonQueryValidator := JsonQueryValidator{
		metadata:       &metadata,
		queryTokenizer: jsonQueryTokenizerMock,
	}

	// Act
	err := jsonQueryValidator.Validate(query)

	// Assert
	assert.Error(t, err)
}

func TestJsonQueryValidator_Validate_ShouldReturnError_WhenQueryTokenLengthIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[]": {},
		"[\"Value\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		"[\"Value\", \"Equal\", \"Test\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
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

	for query, tokens := range queryMap {
		// Arrange
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryValidator := JsonQueryValidator{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		err := jsonQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestJsonQueryValidator_Validate_ShouldReturnError_WhenQueryTokenLengthIsThreeAndThereIsNoneToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Value1\", \"Equal\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "Value1",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		"[\"Value\", \"Equal\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		"[\"Value\", \"Equal1\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "Equal1",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
				Value: "Test",
			},
		},
		"[\"Value\", \"Equal\", \"Test1\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal1",
			},
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "Test1",
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
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryValidator := JsonQueryValidator{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		err := jsonQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestJsonQueryValidator_Validate_ShouldReturnNil_WhenQueryTokenLengthIsThreeAndValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Value\", \"Equal\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeStringValue,
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

	for query, tokens := range queryMap {
		// Arrange
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryValidator := JsonQueryValidator{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		err := jsonQueryValidator.Validate(query)

		// Assert
		assert.NoError(t, err)
	}
}

func TestJsonQueryValidator_Validate_ShouldReturnError_WhenQueryTokenLengthIsTwoAndThereIsNoneToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"O\", [[\"Value\", \"Equal\", \"Test1\"], [\"Value\", \"Equal\", \"Test2\"]]]": {
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "O",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test1",
					},
				},
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test2",
					},
				},
			},
		},
		"[\"Or\", [[\"Value1\", \"Equal\", \"Test1\"], [\"Value\", \"Equal\", \"Test2\"]]]": {
			models.Token{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeNone,
						Value: "Value1",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test1",
					},
				},
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test2",
					},
				},
			},
		},
		"[\"Or\", [[\"Value\", \"Equals\", \"Test1\"], [\"Value\", \"Equal\", \"Test2\"]]]": {
			models.Token{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeNone,
						Value: "Equals",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test1",
					},
				},
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test2",
					},
				},
			},
		},
		"[\"Or\", [[\"Value\", \"Equal\", \"Test1\"], [\"Value\", \"Equal\", \"Test3\"]]]": {
			models.Token{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test1",
					},
				},
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeNone,
						Value: "Test3",
					},
				},
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
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryValidator := JsonQueryValidator{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		err := jsonQueryValidator.Validate(query)

		// Assert
		assert.Error(t, err)
	}
}

func TestJsonQueryValidator_Validate_ShouldReturnNil_WhenQueryTokenLengthIsTwoAndValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Or\", [[\"Value\", \"Equal\", \"Test1\"], [\"Value\", \"Equal\", \"Test2\"]]]": {
			models.Token{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test1",
					},
				},
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
						Value: "Value",
					},
					models.Token{
						Type:  constants.TokenTypeEqual,
						Value: "Equal",
					},
					models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: "Test2",
					},
				},
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
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryValidator := JsonQueryValidator{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		err := jsonQueryValidator.Validate(query)

		// Assert
		assert.NoError(t, err)
	}
}
