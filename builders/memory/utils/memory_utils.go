package utils

import (
	"reflect"
	"strings"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/utils"
)

func CheckEquality(fieldType constants.FieldType, fieldValue interface{}, value interface{}) bool {
	switch fieldType {
	case constants.FieldTypeString:
		castedValue, castedValueErr := utils.String(value)
		castedFieldValue, castedFieldValueErr := utils.String(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && strings.ToLower(castedValue) == strings.ToLower(castedFieldValue) {
			return true
		}
	case constants.FieldTypeStringArray:
		castedValue, castedValueErr := utils.String(value)
		castedFieldValue, castedFieldValueErr := utils.String(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && strings.ToLower(castedValue) == strings.ToLower(castedFieldValue) {
			return true
		}
	case constants.FieldTypeNumber:
		castedValue, castedValueErr := utils.Number(value)
		castedFieldValue, castedFieldValueErr := utils.Number(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue == castedFieldValue {
			return true
		}
	case constants.FieldTypeNumberArray:
		castedValue, castedValueErr := utils.Number(value)
		castedFieldValue, castedFieldValueErr := utils.Number(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue == castedFieldValue {
			return true
		}
	case constants.FieldTypeBoolean:
		castedValue, castedValueErr := utils.Boolean(value)
		castedFieldValue, castedFieldValueErr := utils.Boolean(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue == castedFieldValue {
			return true
		}
	case constants.FieldTypeBooleanArray:
		castedValue, castedValueErr := utils.Boolean(value)
		castedFieldValue, castedFieldValueErr := utils.Boolean(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue == castedFieldValue {
			return true
		}
	case constants.FieldTypeDate:
		castedValue, castedValueErr := utils.Date(value)
		castedFieldValue, castedFieldValueErr := utils.Date(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue.UnixNano() == castedFieldValue.UnixNano() {
			return true
		}
	case constants.FieldTypeDateArray:
		castedValue, castedValueErr := utils.Date(value)
		castedFieldValue, castedFieldValueErr := utils.Date(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue.UnixNano() == castedFieldValue.UnixNano() {
			return true
		}
	case constants.FieldTypeTime:
		castedValue, castedValueErr := utils.Time(value)
		castedFieldValue, castedFieldValueErr := utils.Time(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && *castedValue == *castedFieldValue {
			return true
		}
	case constants.FieldTypeTimeArray:
		castedValue, castedValueErr := utils.Time(value)
		castedFieldValue, castedFieldValueErr := utils.Time(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && *castedValue == *castedFieldValue {
			return true
		}
	case constants.FieldTypeDateTime:
		castedValue, castedValueErr := utils.DateTime(value)
		castedFieldValue, castedFieldValueErr := utils.DateTime(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue.UnixNano() == castedFieldValue.UnixNano() {
			return true
		}
	case constants.FieldTypeDateTimeArray:
		castedValue, castedValueErr := utils.DateTime(value)
		castedFieldValue, castedFieldValueErr := utils.DateTime(fieldValue)
		if castedValueErr == nil && castedFieldValueErr == nil && castedValue.UnixNano() == castedFieldValue.UnixNano() {
			return true
		}
	}

	return false
}

func ObjectToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	value := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typ.Field(i).Name

		if !field.CanInterface() {
			continue
		}

		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				result[fieldName] = nil
			} else {
				result[fieldName] = field.Elem().Interface()
			}
		} else {
			result[fieldName] = field.Interface()
		}
	}

	return result
}
