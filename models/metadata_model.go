package models

import (
	"strings"

	"github.com/filtex/filtex-go/constants"
)

type Metadata struct {
	Fields []Field `json:"fields"`
}

func (m *Metadata) GetFieldType(str string) constants.FieldType {
	for _, v := range m.Fields {
		if strings.ToLower(v.Label) == strings.ToLower(str) || strings.ToLower(v.Name) == strings.ToLower(str) {
			return constants.FieldType(v.Type)
		}
	}

	return constants.FieldTypeUnknown
}

func (m *Metadata) GetFieldName(str string) string {
	for _, v := range m.Fields {
		if strings.ToLower(v.Label) == strings.ToLower(str) || strings.ToLower(v.Name) == strings.ToLower(str) {
			return v.Name
		}
	}

	return str
}

func (m *Metadata) GetFieldValues(str string) []Lookup {
	for _, v := range m.Fields {
		if strings.ToLower(v.Label) == strings.ToLower(str) || strings.ToLower(v.Name) == strings.ToLower(str) {
			return v.Values
		}
	}

	return nil
}
