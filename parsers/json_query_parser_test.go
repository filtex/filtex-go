package parsers

import (
	"errors"
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/expressions"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJsonQueryParser_ShouldReturnError_WhenQueryTokenizerReturnedError(t *testing.T) {
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

	jsonQueryParser := JsonQueryParser{
		metadata:       &metadata,
		queryTokenizer: jsonQueryTokenizerMock,
	}

	// Act
	expression, err := jsonQueryParser.Parse(query)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryTokenLengthIsNotValid(t *testing.T) {
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsTwoAndLogicIsNotToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Test\", [[\"Value\", \"Equal\", \"Test\"], [\"Value\", \"Equal\", \"Test\"]]]": {
			"Test",
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
						Value: "Test",
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
						Value: "Test",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsTwoAndLogicIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Test\", [[\"Value\", \"Equal\", \"Test\"], [\"Value\", \"Equal\", \"Test\"]]]": {
			models.Token{
				Type:  constants.TokenTypeLiteral,
				Value: "Test",
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
						Value: "Test",
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
						Value: "Test",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsTwoAndSecondItemIsNotArray(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Test\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
			models.Token{
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

	for query, tokens := range queryMap {
		// Arrange
		jsonQueryTokenizerMock := tokenizers.NewJsonQueryTokenizerMock()

		jsonQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(tokens, nil)

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnLogicExpression_WhenQueryLengthIsTwo(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"And\", [[\"Value1\", \"Equal\", \"Test1\"], [\"Value2\", \"Equal\", \"Test2\"]]]": {
			models.Token{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
			[]interface{}{
				[]interface{}{
					models.Token{
						Type:  constants.TokenTypeField,
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
						Value: "Value2",
					},
					models.Token{
						Type:  constants.TokenTypeNotEqual,
						Value: "Not Equal",
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
				Name:  "Value1",
				Type:  constants.FieldTypeString.String(),
				Label: "Value1",
				Operators: []string{
					constants.OperatorEqual.String(),
					constants.OperatorNotEqual.String(),
				},
				Values: nil,
			},
			{
				Name:  "Value2",
				Type:  constants.FieldTypeString.String(),
				Label: "Value2",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		var logicExpression *expressions.LogicExpression

		assert.NotNil(t, expression)
		assert.IsType(t, logicExpression, expression)
		assert.NoError(t, err)

		logicExpression = expression.(*expressions.LogicExpression)

		assert.Equal(t, constants.LogicAnd, logicExpression.Logic)
		assert.NotNil(t, logicExpression.Expressions)
		assert.Len(t, logicExpression.Expressions, 2)

		firstExpression, ok := logicExpression.Expressions[0].(*expressions.OperatorExpression)

		assert.True(t, ok)
		assert.Equal(t, "Value1", firstExpression.Field)
		assert.Equal(t, constants.OperatorEqual, firstExpression.Operator)
		assert.Equal(t, "Test1", firstExpression.Value)

		secondExpression, ok := logicExpression.Expressions[1].(*expressions.OperatorExpression)

		assert.True(t, ok)
		assert.Equal(t, "Value2", secondExpression.Field)
		assert.Equal(t, constants.OperatorNotEqual, secondExpression.Operator)
		assert.Equal(t, "Test2", secondExpression.Value)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsThreeAndFieldIsNotToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Value\", \"Equal\", \"Test\"]": {
			"Value",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsThreeAndOperatorIsNotToken(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Value\", \"Equal\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			"Equal",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsThreeAndValueIsNotToken(t *testing.T) {
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
			"Test",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnError_WhenQueryLengthIsThreeAndOperatorIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]interface{}{
		"[\"Value\", \"Op\", \"Test\"]": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
			models.Token{
				Type:  constants.TokenTypeNone,
				Value: "Op",
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestJsonQueryParser_ShouldReturnOperatorExpression_WhenQueryLengthIsThree(t *testing.T) {
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

		jsonQueryParser := JsonQueryParser{
			metadata:       &metadata,
			queryTokenizer: jsonQueryTokenizerMock,
		}

		// Act
		expression, err := jsonQueryParser.Parse(query)

		// Assert
		var operatorExpression *expressions.OperatorExpression

		assert.NotNil(t, expression)
		assert.IsType(t, operatorExpression, expression)
		assert.NoError(t, err)

		operatorExpression = expression.(*expressions.OperatorExpression)

		assert.Equal(t, "Value", operatorExpression.Field)
		assert.Equal(t, constants.OperatorEqual, operatorExpression.Operator)
		assert.Equal(t, "Test", operatorExpression.Value)
	}
}
