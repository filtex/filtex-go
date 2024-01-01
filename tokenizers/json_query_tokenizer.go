package tokenizers

import (
	"encoding/json"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/utils"
)

type jsonQueryTokenizer struct {
	*BaseQueryTokenizer
}

func NewJsonQueryTokenizer(metadata *models.Metadata) JsonQueryTokenizer {
	return &jsonQueryTokenizer{
		BaseQueryTokenizer: NewBaseQueryTokenizer(metadata),
	}
}

func (t *jsonQueryTokenizer) Tokenize(query string) ([]interface{}, error) {
	var data []interface{}
	err := json.Unmarshal([]byte(query), &data)
	if err != nil {
		return nil, err
	}

	return t.tokenizeInternal(data)
}

func (t *jsonQueryTokenizer) tokenizeInternal(data []interface{}) ([]interface{}, error) {
	if len(data) == 3 {
		fieldString, err := utils.String(data[0])
		if err != nil {
			return nil, err
		}
		fieldMatch := t.findMatch(fieldString)
		fieldToken := t.createToken([]models.Token{}, fieldMatch.tokenType, fieldMatch.value)
		if fieldToken == nil || fieldToken.Type == constants.TokenTypeNone {
			fieldToken = &models.Token{
				Type:  constants.TokenTypeNone,
				Value: fieldString,
			}
		}

		operatorString, err := utils.String(data[1])
		if err != nil {
			return nil, err
		}
		operatorMatch := t.findMatch(operatorString)
		operatorToken := t.createToken([]models.Token{*fieldToken}, operatorMatch.tokenType, operatorMatch.value)
		if operatorToken == nil || operatorToken.Type == constants.TokenTypeNone {
			operatorToken = &models.Token{
				Type:  constants.TokenTypeNone,
				Value: operatorString,
			}
		}

		if utils.IsArray(data[2]) {
			valueTokens := make([]models.Token, 0)
			values, err := utils.Array(data[2])
			if err != nil {
				return nil, err
			}

			for _, value := range values {
				valueString, err := utils.String(value)
				if err != nil {
					return nil, err
				}
				valueMatch := t.findMatch(valueString)

				var valueToken *models.Token

				if valueMatch != nil {
					if len(valueString) == len(valueMatch.value) {
						valueToken = t.createToken([]models.Token{*fieldToken, *operatorToken}, valueMatch.tokenType, valueMatch.value)
					} else if t.metadata.GetFieldType(fieldString) == constants.FieldTypeString {
						valueToken = &models.Token{
							Type:  constants.TokenTypeStringValue,
							Value: valueString,
						}
					}
				} else if operatorToken.Type.IsNotComparerTokenType() && valueString == "" {
					valueToken = &models.Token{
						Type:  constants.TokenTypeValue,
						Value: valueString,
					}
				}

				if valueToken == nil || valueToken.Type == constants.TokenTypeNone {
					valueToken = &models.Token{
						Type:  constants.TokenTypeNone,
						Value: valueString,
					}
				}
				valueTokens = append(valueTokens, *valueToken)
			}

			return []interface{}{
				*fieldToken,
				*operatorToken,
				valueTokens,
			}, nil
		} else {
			valueString, err := utils.String(data[2])
			if err != nil {
				return nil, err
			}
			valueMatch := t.findMatch(valueString)

			var valueToken *models.Token

			if valueMatch != nil {
				if len(valueString) == len(valueMatch.value) {
					valueToken = t.createToken([]models.Token{*fieldToken, *operatorToken}, valueMatch.tokenType, valueMatch.value)
				} else if t.metadata.GetFieldType(fieldString) == constants.FieldTypeString {
					valueToken = &models.Token{
						Type:  constants.TokenTypeStringValue,
						Value: valueString,
					}
				}
			} else if operatorToken.Type.IsNotComparerTokenType() && valueString == "" {
				valueToken = &models.Token{
					Type:  constants.TokenTypeValue,
					Value: valueString,
				}
			}

			if valueToken == nil || valueToken.Type == constants.TokenTypeNone {
				valueToken = &models.Token{
					Type:  constants.TokenTypeNone,
					Value: valueString,
				}
			}

			return []interface{}{
				*fieldToken,
				*operatorToken,
				*valueToken,
			}, nil
		}
	}

	if len(data) == 2 {
		logicString, err := utils.String(data[0])
		if err != nil {
			return nil, err
		}
		logicMatch := t.findMatch(logicString)

		logicToken := models.Token{
			Type:  constants.TokenTypeNone,
			Value: logicString,
		}

		logic := constants.ParseLogic(logicString)
		if logic != constants.LogicUnknown {
			logicToken = models.Token{
				Type:  logic.ToTokenType(),
				Value: logicMatch.value,
			}
		}

		if !utils.IsArray(data[1]) {
			return []interface{}{
				logicToken,
				models.Token{
					Type:  constants.TokenTypeNone,
					Value: data[1],
				},
			}, nil
		}

		expressions := make([]interface{}, 0)

		values, err := utils.Array(data[1])
		if err != nil {
			return nil, err
		}

		for _, v := range values {
			ex, err := t.tokenizeInternal(v.([]interface{}))
			if err != nil {
				return nil, err
			}

			expressions = append(expressions, ex)
		}

		return []interface{}{
			logicToken,
			expressions,
		}, nil
	}

	return nil, errors.NewCouldNotBeTokenizedError()
}
