package logics

import (
	"testing"

	"github.com/filtex/filtex-go/builders/postgres/operators"
	"github.com/filtex/filtex-go/builders/postgres/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestAndExpression_ShouldReturnExpression_WhenThereAreExpressions(t *testing.T) {
	// Arrange
	value := "Filtex"

	// Act
	expression := AndLogic{}.Build([]types.PostgresExpression{
		*operators.EqualOperator{}.Build(constants.FieldTypeString, "Value", value, 0),
	})

	// Assert
	assert.NotNil(t, expression)
	assert.NotEmpty(t, expression.Condition)
	assert.Len(t, expression.Args, 1)
	assert.Equal(t, value, expression.Args[0])
}
