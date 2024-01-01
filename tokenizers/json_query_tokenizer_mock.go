package tokenizers

import (
	"github.com/stretchr/testify/mock"
)

type jsonQueryTokenizerMock struct {
	mock.Mock
}

func (j *jsonQueryTokenizerMock) Tokenize(query string) ([]interface{}, error) {
	args := j.Called(query)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]interface{}), args.Error(1)
}

func NewJsonQueryTokenizerMock() *jsonQueryTokenizerMock {
	return &jsonQueryTokenizerMock{}
}
