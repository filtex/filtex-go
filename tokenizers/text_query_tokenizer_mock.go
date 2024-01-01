package tokenizers

import (
	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/mock"
)

type textQueryTokenizerMock struct {
	mock.Mock
}

func (t *textQueryTokenizerMock) Tokenize(text string) (*[]models.Token, error) {
	args := t.Called(text)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*[]models.Token), args.Error(1)
}

func NewTextQueryTokenizerMock() *textQueryTokenizerMock {
	return &textQueryTokenizerMock{}
}
