package types

type MemoryExpression struct {
	Fn func(map[string]interface{}) bool
}
