package operators

import (
	"fmt"

	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
)

type EndWithOperator struct{}

func (EndWithOperator) Build(fieldType constants.FieldType, field string, value interface{}, index int) *types.PostgresExpression {
	if fieldType != constants.FieldTypeString {
		return nil
	}

	return &types.PostgresExpression{
		Condition: fmt.Sprintf("%s ILIKE '%%' || $%v", field, index),
		Args:      []interface{}{value},
	}
}
