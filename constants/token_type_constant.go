package constants

import (
	"github.com/filtex/filtex-go/utils"
)

type TokenType string

const (
	TokenTypeNone               TokenType = ""
	TokenTypeOpenBracket        TokenType = "open-bracket"
	TokenTypeCloseBracket       TokenType = "close-bracket"
	TokenTypeAnd                TokenType = "and"
	TokenTypeOr                 TokenType = "or"
	TokenTypeField              TokenType = "field"
	TokenTypeValue              TokenType = "value"
	TokenTypeEqual              TokenType = "equal"
	TokenTypeNotEqual           TokenType = "not-equal"
	TokenTypeGreaterThan        TokenType = "greater-than"
	TokenTypeGreaterThanOrEqual TokenType = "greater-than-or-equal"
	TokenTypeLessThan           TokenType = "less-than"
	TokenTypeLessThanOrEqual    TokenType = "less-than-or-equal"
	TokenTypeBlank              TokenType = "blank"
	TokenTypeNotBlank           TokenType = "not-blank"
	TokenTypeContain            TokenType = "contain"
	TokenTypeNotContain         TokenType = "not-contain"
	TokenTypeStartWith          TokenType = "start-with"
	TokenTypeNotStartWith       TokenType = "not-start-with"
	TokenTypeEndWith            TokenType = "end-with"
	TokenTypeNotEndWith         TokenType = "not-end-with"
	TokenTypeIn                 TokenType = "in"
	TokenTypeNotIn              TokenType = "not-in"
	TokenTypeComma              TokenType = "comma"
	TokenTypeSlash              TokenType = "slash"
	TokenTypeStringValue        TokenType = "string-value"
	TokenTypeNumberValue        TokenType = "number-value"
	TokenTypeBooleanValue       TokenType = "boolean-value"
	TokenTypeDateValue          TokenType = "date-value"
	TokenTypeTimeValue          TokenType = "time-value"
	TokenTypeDateTimeValue      TokenType = "datetime-value"
	TokenTypeLiteral            TokenType = "literal"
	TokenTypeSpace              TokenType = "space"
)

func (t TokenType) IsFieldTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeField,
	})
}

func (t TokenType) IsOperatorTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeEqual,
		TokenTypeNotEqual,
		TokenTypeBlank,
		TokenTypeNotBlank,
		TokenTypeLessThan,
		TokenTypeLessThanOrEqual,
		TokenTypeGreaterThan,
		TokenTypeGreaterThanOrEqual,
		TokenTypeContain,
		TokenTypeNotContain,
		TokenTypeStartWith,
		TokenTypeNotStartWith,
		TokenTypeEndWith,
		TokenTypeNotEndWith,
		TokenTypeIn,
		TokenTypeNotIn,
	})
}

func (t TokenType) IsComparerTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeEqual,
		TokenTypeNotEqual,
		TokenTypeLessThan,
		TokenTypeLessThanOrEqual,
		TokenTypeGreaterThan,
		TokenTypeGreaterThanOrEqual,
		TokenTypeContain,
		TokenTypeNotContain,
		TokenTypeStartWith,
		TokenTypeNotStartWith,
		TokenTypeEndWith,
		TokenTypeNotEndWith,
		TokenTypeIn,
		TokenTypeNotIn,
	})
}

func (t TokenType) IsNotComparerTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeBlank,
		TokenTypeNotBlank,
	})
}

func (t TokenType) IsSeparatorTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeComma,
		TokenTypeSlash,
	})
}

func (t TokenType) IsPreFieldTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeOpenBracket,
		TokenTypeAnd,
		TokenTypeOr,
	})
}

func (t TokenType) IsValueTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeValue,
		TokenTypeStringValue,
		TokenTypeNumberValue,
		TokenTypeBooleanValue,
		TokenTypeDateValue,
		TokenTypeTimeValue,
		TokenTypeDateTimeValue,
	})
}

func (t TokenType) IsLogicTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeAnd,
		TokenTypeOr,
	})
}

func (t TokenType) IsOpenGroupTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeOpenBracket,
	})
}

func (t TokenType) IsCloseGroupTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeCloseBracket,
	})
}

func (t TokenType) IsMultiAllowedTokenType() bool {
	return utils.IsInAny(t, []TokenType{
		TokenTypeIn,
		TokenTypeNotIn,
	})
}
