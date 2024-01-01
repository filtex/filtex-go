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

func TestTextQueryParser_ShouldReturnError_WhenQueryTokenizerReturnedError(t *testing.T) {
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
	textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()
	textQueryTokenizerMock.
		On("Tokenize", mock.Anything).
		Return(nil, errors.New("some error"))

	textQueryParser := TextQueryParser{
		metadata:       &metadata,
		queryTokenizer: textQueryTokenizerMock,
	}

	// Act
	expression, err := textQueryParser.Parse(query)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestTextQueryParser_ShouldReturnError_WhenQueryTokenizerReturnedEmptyResult(t *testing.T) {
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
	textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()
	textQueryTokenizerMock.
		On("Tokenize", mock.Anything).
		Return(&[]models.Token{}, nil)

	textQueryParser := TextQueryParser{
		metadata:       &metadata,
		queryTokenizer: textQueryTokenizerMock,
	}

	// Act
	expression, err := textQueryParser.Parse(query)

	// Assert
	assert.Nil(t, expression)
	assert.Error(t, err)
}

func TestTextQueryParser_ShouldReturnError_WhenQueryTokenLengthIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"": {},
		"Value": {
			models.Token{
				Type:  constants.TokenTypeField,
				Value: "Value",
			},
		},
		"Value Equal Test Test": {
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
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryParser := TextQueryParser{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		expression, err := textQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestTextQueryParser_ShouldReturnError_WhenLogicIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value Equal Test Xor Value Equal Test": {
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
				Type:  constants.TokenTypeNone,
				Value: "Xor",
			},
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
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryParser := TextQueryParser{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		expression, err := textQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestTextQueryParser_ShouldReturnError_WhenOperatorIsNotValid(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value Op Test": {
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
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryParser := TextQueryParser{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		expression, err := textQueryParser.Parse(query)

		// Assert
		assert.Nil(t, expression)
		assert.Error(t, err)
	}
}

func TestTextQueryParser_ShouldReturnLogicExpression_WhenQueryHasLogic(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value1 Equal Test1 And Value2 Equal Test2": {
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
			models.Token{
				Type:  constants.TokenTypeAnd,
				Value: "And",
			},
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
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryParser := TextQueryParser{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		expression, err := textQueryParser.Parse(query)

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

func TestTextQueryParser_ShouldReturnOperatorExpression_WhenQueryDoesNotHaveLogic(t *testing.T) {
	// Arrange
	queryMap := map[string][]models.Token{
		"Value Equal Test": {
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
		textQueryTokenizerMock := tokenizers.NewTextQueryTokenizerMock()

		textQueryTokenizerMock.
			On("Tokenize", mock.MatchedBy(func(q string) bool { return q == query })).
			Return(&tokens, nil)

		textQueryParser := TextQueryParser{
			metadata:       &metadata,
			queryTokenizer: textQueryTokenizerMock,
		}

		// Act
		expression, err := textQueryParser.Parse(query)

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
