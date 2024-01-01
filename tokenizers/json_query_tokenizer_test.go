package tokenizers

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/assert"
)

func TestJsonQueryTokenizer_Tokenize_ShouldReturnError_WhenQueryIsNotJson(t *testing.T) {
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

	jsonQueryTokenizer := NewJsonQueryTokenizer(&metadata)

	// Act
	result, err := jsonQueryTokenizer.Tokenize(query)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestJsonQueryTokenizer_Tokenize_ShouldReturnError_WhenQueryLengthIsNotValid(t *testing.T) {
	// Arrange
	queries := []string{
		"[]",
		"[\"Value\"]",
		"[\"Value\", \"Equal\", \"Filtex\", \"Go\"]",
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

	jsonQueryTokenizer := NewJsonQueryTokenizer(&metadata)

	for _, query := range queries {
		// Act
		result, err := jsonQueryTokenizer.Tokenize(query)

		// Assert
		assert.Nil(t, result)
		assert.Error(t, err)
	}
}

func TestJsonQueryTokenizer_Tokenize_ShouldReturnTokens_WhenQueryLengthIsThree(t *testing.T) {
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
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		"[\"Value\", \"Equal\", 100]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeNumberValue,
				Value: "100",
			},
		},
		"[\"Value\", \"Equal\", true]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeBooleanValue,
				Value: "true",
			},
		},
		"[\"Value\", \"Equal\", \"2020-01-01\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeDateValue,
				Value: "2020-01-01",
			},
		},
		"[\"Value\", \"Equal\", \"10:12:14\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeTimeValue,
				Value: "10:12:14",
			},
		},
		"[\"Value\", \"Equal\", \"2020-01-01 11:12:13\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeDateTimeValue,
				Value: "2020-01-01 11:12:13",
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

	jsonQueryTokenizer := NewJsonQueryTokenizer(&metadata)

	for query, tokens := range queryMap {
		// Act
		result, err := jsonQueryTokenizer.Tokenize(query)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, tokens[0], result[0])
		assert.Equal(t, tokens[1], result[1])
		assert.Equal(t, tokens[2], result[2])
	}
}

func TestJsonQueryTokenizer_Tokenize_ShouldReturnTokens_WhenQueryLengthIsTwo(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Or\", [[\"Value\", \"Equal\", 100], [\"Value\", \"Equal\", 200]]]": {
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
						Type:  constants.TokenTypeNumberValue,
						Value: "100",
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
						Type:  constants.TokenTypeNumberValue,
						Value: "200",
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

	jsonQueryTokenizer := NewJsonQueryTokenizer(&metadata)

	for query, tokens := range queryMap {
		// Act
		result, err := jsonQueryTokenizer.Tokenize(query)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, tokens[0], result[0])

		tokenExpressions := tokens[1].([]interface{})
		resultExpressions := result[1].([]interface{})

		assert.Equal(t, tokenExpressions[0], resultExpressions[0])
		assert.Equal(t, tokenExpressions[1], resultExpressions[1])
	}
}
