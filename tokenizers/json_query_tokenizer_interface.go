package tokenizers

type JsonQueryTokenizer interface {
	Tokenize(query string) ([]interface{}, error)
}
