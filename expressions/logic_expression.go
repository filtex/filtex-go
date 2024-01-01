package expressions

import (
	"github.com/filtex/filtex-go/constants"
)

type LogicExpression struct {
	Logic       constants.Logic
	Expressions []Expression
}

func NewLogicExpression(logic constants.Logic, expressions []Expression) Expression {
	return &LogicExpression{
		Logic:       logic,
		Expressions: expressions,
	}
}
