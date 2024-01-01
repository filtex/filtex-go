package constants

import (
	"strings"
)

type Logic string

const (
	LogicUnknown Logic = ""
	LogicAnd     Logic = "and"
	LogicOr      Logic = "or"
)

func (l Logic) ToTokenType() TokenType {
	switch l {
	case LogicAnd:
		return TokenTypeAnd
	case LogicOr:
		return TokenTypeOr
	}

	return TokenTypeNone
}

func ParseLogic(str string) Logic {
	str = strings.ToLower(str)
	switch str {
	case string(LogicAnd):
		return LogicAnd
	case string(LogicOr):
		return LogicOr
	default:
		return LogicUnknown
	}
}
