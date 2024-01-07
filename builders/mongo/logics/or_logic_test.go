package logics

import (
	"testing"

	"github.com/filtex/filtex-go/builders/mongo/operators"
	"github.com/filtex/filtex-go/builders/mongo/types"
	"github.com/filtex/filtex-go/constants"
	"github.com/stretchr/testify/assert"
)

func TestOrExpression_ShouldReturnExpression_WhenThereAreExpressions(t *testing.T) {
	// Arrange
	// Act
	expression := OrLogic{}.Build([]*types.MongoExpression{
		operators.EqualOperator{}.Build(constants.FieldTypeString, "Value", "Filtex"),
	})

	// Assert
	assert.NotNil(t, expression)

	value, ok := expression.Condition["$or"]
	assert.True(t, ok)
	assert.Len(t, value, 1)
}
