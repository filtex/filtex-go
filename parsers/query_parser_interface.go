package parsers

import (
	"github.com/filtex/filtex-go/expressions"
)

type QueryParser interface {
	Parse(query string) (expressions.Expression, error)
}
