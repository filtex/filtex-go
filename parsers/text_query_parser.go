package parsers

import (
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/filtex/filtex-go/utils"
)

type TextQueryParser struct {
	metadata       *models.Metadata
	queryTokenizer tokenizers.TextQueryTokenizer
}

func NewTextQueryParser(metadata *models.Metadata, queryTokenizers tokenizers.TextQueryTokenizer) QueryParser {
	return &TextQueryParser{
		metadata:       metadata,
		queryTokenizer: queryTokenizers,
	}
}

func (p *TextQueryParser) Parse(query string) (expressions.Expression, error) {
	tokens, err := p.queryTokenizer.Tokenize(query)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)
	parsed := p.parseTokens(tokens, result, false)
	return p.parseExpression(parsed)
}

func (p *TextQueryParser) parseTokens(queue *[]models.Token, result []interface{}, isValueExpected bool) []interface{} {
	for len(*queue) > 0 {
		token := (*queue)[0]
		*queue = (*queue)[1:]

		if token.Type == constants.TokenTypeSpace {
			continue
		}

		if token.Type.IsFieldTokenType() {
			result = append(result, token)
		} else if token.Type.IsComparerTokenType() {
			result = append(result, token)
		} else if token.Type.IsNotComparerTokenType() {
			result = append(result, token)
			result = append(result, models.Token{
				Type:  constants.TokenTypeValue,
				Value: "",
			})

			if isValueExpected {
				return result
			}
		} else if token.Type.IsValueTokenType() {
			if len(result) > 2 && utils.IsArray(result[2]) {
				inner := result[2].([]interface{})
				inner = append(inner, token)
				result[2] = inner
			} else {
				result = append(result, token)
			}

			if isValueExpected {
				return result
			}
		} else if token.Type.IsLogicTokenType() {
			logicInner := make([]interface{}, 0)
			logicResult := p.parseTokens(queue, logicInner, true)
			newResult := make([]interface{}, 0)
			newResult = append(newResult, token)
			newResult = append(newResult, []interface{}{
				result,
				logicResult,
			})
			result = newResult
		} else if token.Type.IsSeparatorTokenType() {
			if len(result) > 2 && utils.IsArray(result[2]) {
				inner := result[2].([]interface{})
				newResult := make([]interface{}, 0)
				newResult = append(newResult, inner...)
				result[2] = newResult
			} else {
				inner := make([]interface{}, 0)
				inner = append(inner, result[2])
				result[2] = inner
			}
		} else if token.Type.IsOpenGroupTokenType() {
			bracketInner := make([]interface{}, 0)
			result = p.parseTokens(queue, bracketInner, false)
		} else if token.Type.IsCloseGroupTokenType() {
			return result
		} else {
			result = append(result, token)
		}
	}

	return result
}

func (p *TextQueryParser) parseExpression(data []interface{}) (expressions.Expression, error) {
	if len(data) == 3 {
		fieldToken := data[0].(models.Token)
		operatorToken := data[1].(models.Token)

		var value interface{}

		if utils.IsArray(data[2]) {
			v := make([]interface{}, 0)
			for _, valueToken := range data[2].([]interface{}) {
				v = append(v, valueToken.(models.Token).Value)
			}
			value = v
		} else {
			value = data[2].(models.Token).Value
		}

		operator := constants.ParseOperator(string(operatorToken.Type))
		if operator == constants.OperatorUnknown {
			return nil, errors.NewOperatorCouldNotBeParsedError()
		}

		return expressions.NewOperatorExpression(
			p.metadata.GetFieldType(fieldToken.Value.(string)),
			p.metadata.GetFieldName(fieldToken.Value.(string)),
			operator,
			value), nil
	}

	if len(data) == 2 {
		logicToken := data[0].(models.Token)
		expressionList := make([]expressions.Expression, 0)

		logic := constants.ParseLogic(logicToken.Value.(string))
		if logic == constants.LogicUnknown {
			return nil, errors.NewLogicCouldNotBeParsedError()
		}

		for _, v := range data[1].([]interface{}) {
			ex, err := p.parseExpression(v.([]interface{}))
			if err != nil {
				return nil, err
			}

			expressionList = append(expressionList, ex)
		}

		return expressions.NewLogicExpression(logic, expressionList), nil
	}

	return nil, errors.NewCouldNotBeParsedError()
}
