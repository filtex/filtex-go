package options

import (
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/models"
)

type LookupOption struct {
	key    string
	values []models.Lookup
}

func NewLookupOption() *LookupOption {
	return &LookupOption{
		values: make([]models.Lookup, 0),
	}
}

func (l *LookupOption) Key(key string) *LookupOption {
	l.key = key
	return l
}

func (l *LookupOption) Values(values []models.Lookup) *LookupOption {
	l.values = values
	return l
}

func (l *LookupOption) Build() (map[string][]models.Lookup, error) {
	if l.key == "" {
		return nil, errors.NewInvalidLookupKeyError()
	}

	if l.values == nil || len(l.values) == 0 {
		return nil, errors.NewInvalidLookupValuesError()
	}

	return map[string][]models.Lookup{
		l.key: l.values,
	}, nil
}
