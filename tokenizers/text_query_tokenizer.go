package tokenizers

import (
	"regexp"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
)

type textQueryTokenizer struct {
	*BaseQueryTokenizer
}

func NewTextQueryTokenizer(metadata *models.Metadata) TextQueryTokenizer {
	return &textQueryTokenizer{
		BaseQueryTokenizer: NewBaseQueryTokenizer(metadata),
	}
}

func (t *textQueryTokenizer) Tokenize(text string) (*[]models.Token, error) {
	tokens := make([]models.Token, 0)

	remainingText := text

	for len(remainingText) > 0 {
		match := t.findMatch(remainingText)
		if match != nil {
			token := t.createToken(tokens, match.tokenType, match.value)
			if token != nil {
				tokens = append(tokens, *token)
			}
			remainingText = match.remainingText
		} else {
			re := regexp.MustCompile(`^\s+`)
			wsMatch := re.Find([]byte(remainingText))
			if len(wsMatch) > 0 {
				token := t.createToken(tokens, constants.TokenTypeSpace, " ")
				if token != nil {
					tokens = append(tokens, *token)
				}
				remainingText = remainingText[1:]
			} else {
				re := regexp.MustCompile(`(^\S+\s)|^\S+`)
				invalidTokenMatch := re.Find([]byte(remainingText))
				if len(invalidTokenMatch) == 0 {
					break
				}

				token := t.createToken(tokens, constants.TokenTypeNone, string(invalidTokenMatch))
				if token != nil {
					tokens = append(tokens, *token)
				}
				remainingText = remainingText[len(invalidTokenMatch):]
			}
		}
	}

	return &tokens, nil
}
