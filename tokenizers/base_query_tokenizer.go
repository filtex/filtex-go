package tokenizers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/filtex/filtex-go/utils"
)

type BaseQueryTokenizer struct {
	metadata      *models.Metadata
	tokenPatterns []tokenPattern
}

type tokenPattern struct {
	tokenPattern string
	tokenType    constants.TokenType
}

type tokenMatch struct {
	remainingText string
	tokenType     constants.TokenType
	value         string
}

func NewBaseQueryTokenizer(metadata *models.Metadata) *BaseQueryTokenizer {
	tokenizer := BaseQueryTokenizer{
		metadata: metadata,
	}

	tokenizer.tokenPatterns = []tokenPattern{
		{`(?i)^\(`, constants.TokenTypeOpenBracket},
		{`(?i)^\)`, constants.TokenTypeCloseBracket},

		{`(?i)^,`, constants.TokenTypeComma},
		{`(?i)^/`, constants.TokenTypeSlash},

		{`(?i)^and\b`, constants.TokenTypeAnd},
		{`(?i)^&&`, constants.TokenTypeAnd},
		{`(?i)^or\b`, constants.TokenTypeOr},
		{`(?i)^\|\|`, constants.TokenTypeOr},

		{`(?i)^=`, constants.TokenTypeEqual},
		{`(?i)^equal\b`, constants.TokenTypeEqual},
		{`(?i)^!=`, constants.TokenTypeNotEqual},
		{`(?i)^not equal\b`, constants.TokenTypeNotEqual},

		{`(?i)^>=`, constants.TokenTypeGreaterThanOrEqual},
		{`(?i)^greater than or equal\b`, constants.TokenTypeGreaterThanOrEqual},
		{`(?i)^>`, constants.TokenTypeGreaterThan},
		{`(?i)^greater than\b`, constants.TokenTypeGreaterThan},

		{`(?i)^<=`, constants.TokenTypeLessThanOrEqual},
		{`(?i)^less than or equal\b`, constants.TokenTypeLessThanOrEqual},
		{`(?i)^<`, constants.TokenTypeLessThan},
		{`(?i)^less than\b`, constants.TokenTypeLessThan},

		{`(?i)^\[\]`, constants.TokenTypeBlank},
		{`(?i)^blank\b`, constants.TokenTypeBlank},
		{`(?i)^!\[\]`, constants.TokenTypeNotBlank},
		{`(?i)^not blank\b`, constants.TokenTypeNotBlank},

		{`(?i)^~\*`, constants.TokenTypeStartWith},
		{`(?i)^start with\b`, constants.TokenTypeStartWith},
		{`(?i)^!~\*`, constants.TokenTypeNotStartWith},
		{`(?i)^not start with\b`, constants.TokenTypeNotStartWith},

		{`(?i)^\*~`, constants.TokenTypeEndWith},
		{`(?i)^end with\b`, constants.TokenTypeEndWith},
		{`(?i)^!\*~`, constants.TokenTypeNotEndWith},
		{`(?i)^not end with\b`, constants.TokenTypeNotEndWith},

		{`(?i)^~`, constants.TokenTypeContain},
		{`(?i)^contain\b`, constants.TokenTypeContain},
		{`(?i)^!~`, constants.TokenTypeNotContain},
		{`(?i)^not contain\b`, constants.TokenTypeNotContain},

		{`(?i)^in\b`, constants.TokenTypeIn},
		{`(?i)^not in\b`, constants.TokenTypeNotIn},
	}

	for _, field := range metadata.Fields {
		tokenizer.tokenPatterns = append(tokenizer.tokenPatterns, tokenPattern{
			tokenPattern: fmt.Sprintf(`(?i)^%s\b`, field.Label),
			tokenType:    constants.TokenTypeField,
		})
		tokenizer.tokenPatterns = append(tokenizer.tokenPatterns, tokenPattern{
			tokenPattern: fmt.Sprintf(`(?i)^%s\b`, field.Name),
			tokenType:    constants.TokenTypeField,
		})
	}

	tokenizer.tokenPatterns = append(tokenizer.tokenPatterns,
		tokenPattern{`(?i)^"[^"]*"`, constants.TokenTypeStringValue},
		tokenPattern{`(?i)^\'[^\']*\'`, constants.TokenTypeStringValue},
		tokenPattern{`(?i)^\d\d\d\d-\d\d-\d\d \d\d:\d\d(:\d\d)?`, constants.TokenTypeDateTimeValue},
		tokenPattern{`(?i)^\d\d\d\d-\d\d-\d\d`, constants.TokenTypeDateValue},
		tokenPattern{`(?i)^\d\d:\d\d(:\d\d)?`, constants.TokenTypeTimeValue},
		tokenPattern{`(?i)^(\d+h)?( ?\d+m)?( ?\d+s)?`, constants.TokenTypeTimeValue},
		tokenPattern{`(?i)^[0-9]+([.][0-9]+)?`, constants.TokenTypeNumberValue},
		tokenPattern{`(?i)^(true|false)`, constants.TokenTypeBooleanValue},
		tokenPattern{`(?i)^[a-zA-Z0-9-_]+`, constants.TokenTypeLiteral})

	return &tokenizer
}

func (t *BaseQueryTokenizer) createToken(tokens []models.Token, tokenType constants.TokenType, value string) *models.Token {
	if tokenType == constants.TokenTypeSpace {
		if len(tokens) > 0 && tokens[len(tokens)-1].Type == constants.TokenTypeSpace {
			return nil
		}
		return &models.Token{
			Type:  tokenType,
			Value: value,
		}
	}

	allTokens := make([]models.Token, 0)

	var lastToken models.Token
	var lastTokenType constants.TokenType
	var lastFieldToken *models.Token
	var lastOperatorToken *models.Token

	for _, v := range tokens {
		if v.Type == constants.TokenTypeSpace {
			continue
		}

		allTokens = append(allTokens, v)

		if v.Type.IsFieldTokenType() {
			lastFieldToken = &models.Token{
				Type:  v.Type,
				Value: v.Value,
			}
		} else if v.Type.IsOperatorTokenType() {
			lastOperatorToken = &models.Token{
				Type:  v.Type,
				Value: v.Value,
			}
		}

		lastToken = v
		lastTokenType = v.Type
	}

	if len(allTokens) == 0 {
		if tokenType == constants.TokenTypeField || tokenType == constants.TokenTypeLiteral {
			if t.validateField(value) {
				return &models.Token{
					Type:  constants.TokenTypeField,
					Value: value,
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		} else if tokenType.IsOpenGroupTokenType() {
			return &models.Token{
				Type:  tokenType,
				Value: value,
			}
		}
	} else if tokenType == constants.TokenTypeField {
		if lastTokenType.IsPreFieldTokenType() {
			if t.validateField(value) {
				return &models.Token{
					Type:  constants.TokenTypeField,
					Value: value,
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		} else if lastTokenType.IsComparerTokenType() || lastTokenType.IsSeparatorTokenType() {
			lookupValue := value

			for _, v := range t.metadata.GetFieldValues(lastFieldToken.Value.(string)) {
				if strings.ToLower(v.Name) == strings.ToLower(value) {
					lookupValue, _ = utils.String(v.Value)
					break
				}
			}

			if t.validateValue(lastFieldToken.Value, lookupValue) {
				if lastOperatorToken != nil && lastOperatorToken.Type.IsComparerTokenType() {
					return &models.Token{
						Type:  constants.TokenTypeValue,
						Value: t.castValue(lastFieldToken.Value, lookupValue),
					}
				} else {
					return &models.Token{
						Type:  constants.TokenTypeNone,
						Value: value,
					}
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		}
	} else if tokenType == constants.TokenTypeLiteral {
		if lastTokenType.IsComparerTokenType() || lastTokenType.IsSeparatorTokenType() {
			lookupValue := value

			for _, v := range t.metadata.GetFieldValues(lastFieldToken.Value.(string)) {
				if strings.ToLower(v.Name) == strings.ToLower(value) {
					lookupValue, _ = utils.String(v.Value)
					break
				}
			}

			if lastFieldToken != nil && t.validateValue(lastFieldToken.Value, lookupValue) {
				if lastOperatorToken != nil && lastOperatorToken.Type.IsComparerTokenType() {
					return &models.Token{
						Type:  constants.TokenTypeValue,
						Value: t.castValue(lastFieldToken.Value, lookupValue),
					}
				} else {
					return &models.Token{
						Type:  constants.TokenTypeNone,
						Value: value,
					}
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		} else if lastTokenType.IsPreFieldTokenType() {
			if t.validateField(value) {
				return &models.Token{
					Type:  constants.TokenTypeField,
					Value: value,
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		}
	} else if tokenType.IsValueTokenType() {
		if lastTokenType.IsComparerTokenType() || lastTokenType.IsSeparatorTokenType() {
			if lastFieldToken != nil && t.validateValue(lastFieldToken.Value, value) {
				if lastOperatorToken != nil && lastOperatorToken.Type.IsComparerTokenType() {
					return &models.Token{
						Type:  tokenType,
						Value: t.castValue(lastFieldToken.Value, value),
					}
				} else {
					return &models.Token{
						Type:  constants.TokenTypeNone,
						Value: value,
					}
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		}
	} else if tokenType.IsOperatorTokenType() {
		if lastTokenType == constants.TokenTypeField {
			operator := tokenType.ToOperator()

			if t.validateOperator(lastToken.Value, operator.String()) {
				return &models.Token{
					Type:  tokenType,
					Value: value,
				}
			} else {
				return &models.Token{
					Type:  constants.TokenTypeNone,
					Value: value,
				}
			}
		}
	} else if tokenType.IsLogicTokenType() {
		if lastTokenType.IsValueTokenType() || lastTokenType.IsCloseGroupTokenType() || lastTokenType.IsNotComparerTokenType() {
			return &models.Token{
				Type:  tokenType,
				Value: value,
			}
		}
	} else if tokenType.IsOpenGroupTokenType() {
		if lastTokenType.IsLogicTokenType() || lastTokenType.IsOpenGroupTokenType() {
			return &models.Token{
				Type:  tokenType,
				Value: value,
			}
		}
	} else if tokenType.IsCloseGroupTokenType() {
		openGroupTokenCount := 0
		closeGroupTokenCount := 0

		for _, v := range allTokens {
			if v.Type.IsOpenGroupTokenType() {
				openGroupTokenCount++
			} else if v.Type.IsCloseGroupTokenType() {
				closeGroupTokenCount++
			}
		}

		if openGroupTokenCount > closeGroupTokenCount && (lastTokenType.IsValueTokenType() || lastTokenType.IsCloseGroupTokenType() || lastTokenType.IsNotComparerTokenType()) {
			return &models.Token{
				Type:  tokenType,
				Value: value,
			}
		}
	} else if tokenType.IsSeparatorTokenType() {
		if lastOperatorToken != nil && lastOperatorToken.Type.IsComparerTokenType() && lastOperatorToken.Type.IsMultiAllowedTokenType() {
			if lastTokenType.IsValueTokenType() {
				return &models.Token{
					Type:  tokenType,
					Value: value,
				}
			}
		}
	}

	return &models.Token{
		Type:  constants.TokenTypeNone,
		Value: value,
	}
}

func (t *BaseQueryTokenizer) findMatch(text string) *tokenMatch {
	for _, v := range t.tokenPatterns {
		re := regexp.MustCompile(v.tokenPattern)
		match := string(re.Find([]byte(text)))

		if len(match) > 0 {
			remainingText := ""
			if len(match) != len(text) {
				remainingText = text[len(match):]
			}

			return &tokenMatch{
				remainingText: remainingText,
				tokenType:     v.tokenType,
				value:         match,
			}
		}
	}

	return nil
}

func (t *BaseQueryTokenizer) validateField(field interface{}) bool {
	for _, v := range t.metadata.Fields {
		if strings.ToLower(v.Name) == strings.ToLower(field.(string)) || strings.ToLower(v.Label) == strings.ToLower(field.(string)) {
			return true
		}
	}
	return false
}

func (t *BaseQueryTokenizer) validateOperator(field interface{}, operator interface{}) bool {
	var fieldValue *models.Field

	for _, v := range t.metadata.Fields {
		if strings.ToLower(v.Name) == strings.ToLower(field.(string)) || strings.ToLower(v.Label) == strings.ToLower(field.(string)) {
			fieldValue = &models.Field{
				Operators: v.Operators,
			}
		}
	}

	if fieldValue == nil {
		return false
	}

	for _, v := range fieldValue.Operators {
		if strings.ToLower(v) == strings.ToLower(operator.(string)) {
			return true
		}
	}

	return false
}

func (t *BaseQueryTokenizer) validateValue(field interface{}, value string) bool {
	var fieldValue *models.Field

	for _, v := range t.metadata.Fields {
		if strings.ToLower(v.Label) == strings.ToLower(field.(string)) || strings.ToLower(v.Name) == strings.ToLower(field.(string)) {
			fieldValue = &models.Field{
				Type:   v.Type,
				Values: v.Values,
			}
		}
	}

	if fieldValue == nil {
		return false
	}

	if len(fieldValue.Values) == 0 {
		switch fieldValue.Type {
		case constants.FieldTypeString.String():
			return utils.IsString(value)
		case constants.FieldTypeStringArray.String():
			return utils.IsString(value)
		case constants.FieldTypeNumber.String():
			return utils.IsNumber(value)
		case constants.FieldTypeNumberArray.String():
			return utils.IsNumber(value)
		case constants.FieldTypeBoolean.String():
			return utils.IsBoolean(value)
		case constants.FieldTypeBooleanArray.String():
			return utils.IsBoolean(value)
		case constants.FieldTypeDate.String():
			return utils.IsDate(value)
		case constants.FieldTypeDateArray.String():
			return utils.IsDate(value)
		case constants.FieldTypeTime.String():
			return utils.IsTime(value)
		case constants.FieldTypeTimeArray.String():
			return utils.IsTime(value)
		case constants.FieldTypeDateTime.String():
			return utils.IsDateTime(value)
		case constants.FieldTypeDateTimeArray.String():
			return utils.IsDateTime(value)
		}

		return false
	}

	for _, v := range fieldValue.Values {
		vVal, _ := utils.String(v.Value)
		vName, _ := utils.String(v.Name)
		val, _ := utils.String(value)

		if strings.ToLower(vVal) == strings.ToLower(val) || strings.ToLower(vName) == strings.ToLower(val) {
			return true
		}
	}

	return false
}

func (t *BaseQueryTokenizer) castValue(field interface{}, value interface{}) interface{} {
	var fieldType string

	for _, v := range t.metadata.Fields {
		if strings.ToLower(v.Label) == strings.ToLower(field.(string)) {
			fieldType = v.Type
			break
		}
	}

	if fieldType == "" {
		return value
	}

	switch fieldType {
	case constants.FieldTypeString.String():
		s, _ := utils.String(value)
		return strings.Trim(strings.Trim(s, `"`), `'`)
	case constants.FieldTypeStringArray.String():
		s, _ := utils.String(value)
		return strings.Trim(strings.Trim(s, `"`), `'`)
	case constants.FieldTypeNumber.String():
		f, _ := utils.Number(value)
		return f
	case constants.FieldTypeNumberArray.String():
		f, _ := utils.Number(value)
		return f
	case constants.FieldTypeBoolean.String():
		b, _ := utils.Boolean(value)
		return b
	case constants.FieldTypeBooleanArray.String():
		b, _ := utils.Boolean(value)
		return b
	case constants.FieldTypeDate.String():
		t, _ := utils.Date(value)
		return t
	case constants.FieldTypeDateArray.String():
		t, _ := utils.Date(value)
		return t
	case constants.FieldTypeTime.String():
		d, _ := utils.Time(value)
		return d
	case constants.FieldTypeTimeArray.String():
		d, _ := utils.Time(value)
		return d
	case constants.FieldTypeDateTime.String():
		dt, _ := utils.DateTime(value)
		return dt
	case constants.FieldTypeDateTimeArray.String():
		dt, _ := utils.DateTime(value)
		return dt
	}

	return value
}
