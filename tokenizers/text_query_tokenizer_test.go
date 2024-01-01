package tokenizers

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/assert"
)

func TestTextQueryTokenizer_Tokenize_ShouldReturnTokens(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"": {},
		"Value": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		"Value Equal Test": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeValue,
				Value: "Test",
			},
		},
		"Value Equal Test1 Or Value Equal Test2": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeValue,
				Value: "Test1",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeOr,
				Value: "Or",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeEqual,
				Value: "Equal",
			},
			models.Token{
				Type:  constants.TokenTypeSpace,
				Value: " ",
			},
			models.Token{
				Type:  constants.TokenTypeValue,
				Value: "Test2",
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

	textQueryTokenizer := NewTextQueryTokenizer(&metadata)

	for query, tokens := range queryMap {
		// Act
		result, err := textQueryTokenizer.Tokenize(query)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)

		assert.Len(t, *result, len(tokens))

		for i, v := range *result {
			assert.Equal(t, tokens[i].Type, v.Type)
			assert.Equal(t, tokens[i].Value, v.Value)
		}
	}
}
