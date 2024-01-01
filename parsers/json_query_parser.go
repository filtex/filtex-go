package parsers

import (
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/filtex/filtex-go/utils"
)

type JsonQueryParser struct {
	metadata       *models.Metadata
	queryTokenizer tokenizers.JsonQueryTokenizer
}

func NewJsonQueryParser(metadata *models.Metadata, queryTokenizer tokenizers.JsonQueryTokenizer) QueryParser {
	return &JsonQueryParser{
		metadata:       metadata,
		queryTokenizer: queryTokenizer,
	}
}

func (p *JsonQueryParser) Parse(query string) (expressions.Expression, error) {
	tokens, err := p.queryTokenizer.Tokenize(query)
	if err != nil {
		return nil, err
	}

	return p.parseInternal(tokens)
}

func (p *JsonQueryParser) parseInternal(data []interface{}) (expressions.Expression, error) {
	if len(data) == 3 {
		fieldToken, ok := data[0].(models.Token)
		if !ok {
			return nil, errors.NewCouldNotBeParsedError()
		}

		operatorToken, ok := data[1].(models.Token)
		if !ok {
			return nil, errors.NewCouldNotBeParsedError()
		}

		var value interface{}

		if utils.IsArray(data[2]) {
			valueTokens, ok := data[2].([]models.Token)
			if !ok {
				return nil, errors.NewCouldNotBeParsedError()
			}

			v := make([]interface{}, 0)
			for _, valueToken := range valueTokens {
				v = append(v, valueToken.Value)
			}
			value = v
		} else {
			valueToken, ok := data[2].(models.Token)
			if !ok {
				return nil, errors.NewCouldNotBeParsedError()
			}

			value = valueToken.Value
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
		logicToken, ok := data[0].(models.Token)
		if !ok {
			return nil, errors.NewCouldNotBeParsedError()
		}

		expressionList := make([]expressions.Expression, 0)

		logic := constants.ParseLogic(logicToken.Value.(string))
		if logic == constants.LogicUnknown {
			return nil, errors.NewLogicCouldNotBeParsedError()
		}

		expressionTokens, ok := data[1].([]interface{})
		if !ok {
			return nil, errors.NewCouldNotBeParsedError()
		}

		for _, v := range expressionTokens {
			ex, err := p.parseInternal(v.([]interface{}))
			if err != nil {
				return nil, err
			}

			expressionList = append(expressionList, ex)
		}

		return expressions.NewLogicExpression(logic, expressionList), nil
	}

	return nil, errors.NewCouldNotBeParsedError()
}
