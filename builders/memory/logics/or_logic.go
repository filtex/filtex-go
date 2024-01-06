package logics

import (
	"github.com/filtex/filtex-go/builders/memory/types"
)

type OrLogic struct{}

func (OrLogic) Build(expressions []*types.MemoryExpression) *types.MemoryExpression {
	return &types.MemoryExpression{
		Fn: func(data map[string]interface{}) bool {
			for _, v := range expressions {
				if v.Fn(data) {
					return true
				}
			}
			return false
		},
	}
}
