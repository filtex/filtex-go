package filtex

import (
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/options"
)

type Filtex struct {
	metadata *models.Metadata
}

func New(opts ...options.Option) (*Filtex, error) {
	f := &Filtex{}

	lookups := make(map[string][]models.Lookup)

	for _, v := range opts {
		if lookupOption, ok := v.(*options.LookupOption); ok {
			build, err := lookupOption.Build()
			if err != nil {
				return nil, err
			}
			for lk, lv := range build {
				lookups[lk] = lv
			}
		}
	}

	fields := make([]models.Field, 0)

	for _, v := range opts {
		if fieldOption, ok := v.(*options.FieldOption); ok {
			build, err := fieldOption.Build(lookups)
			if err != nil {
				return nil, err
			}
			fields = append(fields, *build)
		}
	}

	f.metadata = &models.Metadata{
		Fields: fields,
	}

	return f, nil
}

func (f *Filtex) Metadata() (*models.Metadata, error) {
	return f.metadata, nil
}
