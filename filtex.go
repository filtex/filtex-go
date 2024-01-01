package filtex

import (
	"github.com/filtex/filtex-go/expressions"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/options"
	"github.com/filtex/filtex-go/parsers"
	"github.com/filtex/filtex-go/tokenizers"
	"github.com/filtex/filtex-go/validators"
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

func (f *Filtex) ExpressionFromJson(query string) (expressions.Expression, error) {
	return parsers.NewJsonQueryParser(f.metadata, tokenizers.NewJsonQueryTokenizer(f.metadata)).Parse(query)
}

func (f *Filtex) ExpressionFromText(query string) (expressions.Expression, error) {
	return parsers.NewTextQueryParser(f.metadata, tokenizers.NewTextQueryTokenizer(f.metadata)).Parse(query)
}

func (f *Filtex) ValidateFromJson(query string) error {
	return validators.NewJsonQueryValidator(f.metadata, tokenizers.NewJsonQueryTokenizer(f.metadata)).Validate(query)
}

func (f *Filtex) ValidateFromText(query string) error {
	return validators.NewTextQueryValidator(f.metadata, tokenizers.NewTextQueryTokenizer(f.metadata)).Validate(query)
}
