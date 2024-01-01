package models

import "github.com/filtex/filtex-go/constants"

type Token struct {
	Type  constants.TokenType
	Value interface{}
}
