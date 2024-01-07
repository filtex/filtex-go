package logics

import (
	"fmt"
	"strings"

	"github.com/filtex/filtex-go/builders/postgres/types"
)

type OrLogic struct{}

func (OrLogic) Build(expressions []types.PostgresExpression) *types.PostgresExpression {
	conditions := make([]string, 0)
	args := make([]interface{}, 0)

	for _, v := range expressions {
		conditions = append(conditions, fmt.Sprintf("(%s)", v.Condition))
		args = append(args, v.Args...)
	}

	return &types.PostgresExpression{
		Condition: strings.Join(conditions, " OR "),
		Args:      args,
	}
}
