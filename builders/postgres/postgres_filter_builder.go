package postgres

import (
	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
)

type PostgresFilterBuilder struct {
}

func NewPostgresFilterBuilder() *PostgresFilterBuilder {
	return &PostgresFilterBuilder{}
}

func (b *PostgresFilterBuilder) Build(ex expressions.Expression) (*types.PostgresExpression, error) {
	return nil, errors.NewCouldNotBeBuiltError()
}
