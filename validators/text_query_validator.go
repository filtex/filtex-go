package validators

import (
	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/tokenizers"
)

type TextQueryValidator struct {
	metadata       *models.Metadata
	queryTokenizer tokenizers.TextQueryTokenizer
}

func NewTextQueryValidator(metadata *models.Metadata, queryTokenizer tokenizers.TextQueryTokenizer) *TextQueryValidator {
	return &TextQueryValidator{
		metadata:       metadata,
		queryTokenizer: queryTokenizer,
	}
}

func (v *TextQueryValidator) Validate(query string) error {
	tokens, err := v.queryTokenizer.Tokenize(query)
	if err != nil {
		return err
	}

	return v.validateInternal(tokens)
}

func (v *TextQueryValidator) validateInternal(tokens *[]models.Token) error {
	tokensExceptSpace := make([]models.Token, 0)
	openGroupTokenCount := 0
	closeGroupTokenCount := 0

	for _, v := range *tokens {
		if v.Type == constants.TokenTypeNone {
			return errors.NewInvalidTokenError()
		}

		if v.Type != constants.TokenTypeSpace {
			tokensExceptSpace = append(tokensExceptSpace, v)
		}

		if v.Type.IsOpenGroupTokenType() {
			openGroupTokenCount++
		}

		if v.Type.IsCloseGroupTokenType() {
			closeGroupTokenCount++
		}
	}

	if len(tokensExceptSpace) == 0 {
		return nil
	}

	lastTokenType := tokensExceptSpace[len(tokensExceptSpace)-1].Type

	if lastTokenType.IsFieldTokenType() ||
		lastTokenType.IsComparerTokenType() ||
		lastTokenType.IsSeparatorTokenType() ||
		lastTokenType.IsLogicTokenType() ||
		lastTokenType.IsOpenGroupTokenType() {
		return errors.NewInvalidLastTokenError()
	}

	if openGroupTokenCount != closeGroupTokenCount {
		return errors.NewMismatchedBracketsError()
	}

	return nil
}
