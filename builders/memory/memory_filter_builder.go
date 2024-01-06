package memory

import (
	"github.com/filtex/filtex-go/builders/memory/types"
	"github.com/filtex/filtex-go/errors"
	"github.com/filtex/filtex-go/expressions"
)

type MemoryFilterBuilder struct {
}

func NewMemoryFilterBuilder() *MemoryFilterBuilder {
	return &MemoryFilterBuilder{}
}

func (b *MemoryFilterBuilder) Build(expression expressions.Expression) (*types.MemoryExpression, error) {
	return nil, errors.NewCouldNotBeBuiltError()
}
