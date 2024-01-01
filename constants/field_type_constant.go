package constants

type FieldType string

const (
	FieldTypeUnknown       FieldType = ""
	FieldTypeString        FieldType = "string"
	FieldTypeNumber        FieldType = "number"
	FieldTypeBoolean       FieldType = "boolean"
	FieldTypeDate          FieldType = "date"
	FieldTypeTime          FieldType = "time"
	FieldTypeDateTime      FieldType = "datetime"
	FieldTypeStringArray   FieldType = "string-array"
	FieldTypeNumberArray   FieldType = "number-array"
	FieldTypeBooleanArray  FieldType = "boolean-array"
	FieldTypeDateArray     FieldType = "date-array"
	FieldTypeTimeArray     FieldType = "time-array"
	FieldTypeDateTimeArray FieldType = "datetime-array"
)

func (f FieldType) String() string {
	return string(f)
}

func (f FieldType) IsArray() bool {
	return f == FieldTypeStringArray ||
		f == FieldTypeNumberArray ||
		f == FieldTypeBooleanArray ||
		f == FieldTypeDateArray ||
		f == FieldTypeTimeArray ||
		f == FieldTypeDateTimeArray
}
