package types

type PostgresExpression struct {
	Condition string
	Args      []interface{}
}
