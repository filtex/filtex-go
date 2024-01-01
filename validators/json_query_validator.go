package validators

import (
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/filtex/filtex-go/utils"
)

type JsonQueryValidator struct {
	metadata       *models.Metadata
	queryTokenizer tokenizers.JsonQueryTokenizer
}

func NewJsonQueryValidator(metadata *models.Metadata, queryTokenizer tokenizers.JsonQueryTokenizer) *JsonQueryValidator {
	return &JsonQueryValidator{
		metadata:       metadata,
		queryTokenizer: queryTokenizer,
	}
}

func (v *JsonQueryValidator) Validate(query string) error {
	tokens, err := v.queryTokenizer.Tokenize(query)
	if err != nil {
		return err
	}

	return v.validateInternal(tokens)
}

func (v *JsonQueryValidator) validateInternal(data []interface{}) error {
	if len(data) == 3 {
		fieldToken := data[0].(models.Token)
		operatorToken := data[1].(models.Token)

		if utils.IsArray(data[2]) {
			for _, valueToken := range data[2].([]models.Token) {
				if valueToken.Type == constants.TokenTypeNone {
					return errors.NewInvalidValueError()
				}
			}
		} else {
			valueToken := data[2].(models.Token)
			if valueToken.Type == constants.TokenTypeNone {
				return errors.NewInvalidValueError()
			}
		}

		if fieldToken.Type == constants.TokenTypeNone {
			return errors.NewInvalidFieldError()
		}

		if operatorToken.Type == constants.TokenTypeNone {
			return errors.NewInvalidOperatorError()
		}
	} else if len(data) == 2 {
		logicToken := data[0].(models.Token)

		if logicToken.Type == constants.TokenTypeNone {
			return errors.NewInvalidLogicError()
		}

		for _, i := range data[1].([]interface{}) {
			err := v.validateInternal(i.([]interface{}))
			if err != nil {
				return err
			}
		}

	} else {
		return errors.NewCouldNotBeValidatedError()
	}

	return nil
}
