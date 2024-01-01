package tokenizers

import "github.com/filtex/filtex-go/models"

type TextQueryTokenizer interface {
	Tokenize(text string) (*[]models.Token, error)
}
