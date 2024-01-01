package options

import (
	"testing"

	"github.com/filtex/filtex-go/constants"
	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/assert"
)

func TestNewFieldOption_ShouldReturnFieldOption(t *testing.T) {
	// Act
	opt := NewFieldOption()

	// Assert
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.name)
	assert.NotNil(t, opt.label)
	assert.NotNil(t, opt.lookup)
	assert.NotNil(t, opt.fieldType)
	assert.False(t, opt.isArray)
	assert.False(t, opt.isNullable)
}

func TestFieldOption_String_ShouldSetFieldTypeAsStringAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.String()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeString, result.fieldType)
}

func TestFieldOption_Number_ShouldSetFieldTypeAsNumberAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Number()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeNumber, result.fieldType)
}

func TestFieldOption_Boolean_ShouldSetFieldTypeAsBooleanAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Boolean()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeBoolean, result.fieldType)
}

func TestFieldOption_Date_ShouldSetFieldTypeAsDateAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Date()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeDate, result.fieldType)
}

func TestFieldOption_Time_ShouldSetFieldTypeAsTimeAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Time()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeTime, result.fieldType)
}

func TestFieldOption_DateTime_ShouldSetFieldTypeAsDateTimeAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.DateTime()

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, constants.FieldTypeDateTime, result.fieldType)
}

func TestFieldOption_Array_ShouldSetIsArrayAsTrueAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Array()

	// Assert
	assert.NotNil(t, result)
	assert.True(t, result.isArray)
}

func TestFieldOption_Nullable_ShouldSetIsNullableAsTrueAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Nullable()

	// Assert
	assert.NotNil(t, result)
	assert.True(t, result.isNullable)
}

func TestFieldOption_Name_ShouldSetNameAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Name("Some Name")

	// Assert
	assert.NotNil(t, result)
	assert.NotNil(t, result.name)
	assert.Equal(t, "Some Name", result.name)
}

func TestFieldOption_Label_ShouldSetLabelAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Label("Some Label")

	// Assert
	assert.NotNil(t, result)
	assert.NotNil(t, result.label)
	assert.Equal(t, "Some Label", result.label)
}

func TestFieldOption_Lookup_ShouldSetLookupAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewFieldOption()

	// Act
	result := opt.Lookup("some_key")

	// Assert
	assert.NotNil(t, result)
	assert.NotNil(t, result.lookup)
	assert.Equal(t, "some_key", result.lookup)
}

func TestFieldOption_Build_ShouldReturnError_WhenFieldTypeIsNotDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFieldOption_Build_ShouldReturnError_WhenNameIsNotDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFieldOption_Build_ShouldReturnError_WhenLabelIsNotDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Name("Some Name")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsStringArray_WhenTypeIsStringAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeStringArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsNumberArray_WhenTypeIsNumberAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Number().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeNumberArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsBooleanArray_WhenTypeIsBooleanAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Boolean().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeBooleanArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsDateArray_WhenTypeIsDateAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Date().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeDateArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsTimeArray_WhenTypeIsTimeAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Time().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeTimeArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetFieldTypeAsDateTimeArray_WhenTypeIsDateTimeAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		DateTime().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeDateTimeArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
}

func TestFieldOption_Build_ShouldSetValues_WhenValuesAreDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		DateTime().
		Array().
		Name("Some Name").
		Label("Some Label").
		Lookup("some_key")

	// Act
	result, err := opt.Build(map[string][]models.Lookup{
		"some_key": {
			{
				Name:  "Enabled",
				Value: true,
			},
			{
				Name:  "Disabled",
				Value: false,
			},
		},
	})

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, constants.FieldTypeDateTimeArray.String(), result.Type)
	assert.Equal(t, "Some Name", result.Name)
	assert.Equal(t, "Some Label", result.Label)
	assert.Len(t, result.Values, 2)
}

func TestFieldOption_Build_ShouldAddDefaultOperators_WhenDefinitionsAreValid(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorNotEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorIn.String())
	assert.Contains(t, result.Operators, constants.OperatorNotIn.String())
}

func TestFieldOption_Build_ShouldAddBlankOperators_WhenArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorBlank.String())
	assert.Contains(t, result.Operators, constants.OperatorNotBlank.String())
}

func TestFieldOption_Build_ShouldAddBlankOperators_WhenNullableIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Nullable().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorBlank.String())
	assert.Contains(t, result.Operators, constants.OperatorNotBlank.String())
}

func TestFieldOption_Build_ShouldAddContainOperators_WhenArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorContain.String())
	assert.Contains(t, result.Operators, constants.OperatorNotContain.String())
}

func TestFieldOption_Build_ShouldAddContainOperators_WhenTypeIsString(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		String().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorContain.String())
	assert.Contains(t, result.Operators, constants.OperatorNotContain.String())
	assert.Contains(t, result.Operators, constants.OperatorStartWith.String())
	assert.Contains(t, result.Operators, constants.OperatorNotStartWith.String())
	assert.Contains(t, result.Operators, constants.OperatorEndWith.String())
	assert.Contains(t, result.Operators, constants.OperatorNotEndWith.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsNumber(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Number().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsNumberAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Number().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsDate(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Date().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsDateAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Date().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsTime(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Time().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsTimeAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		Time().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsDateTime(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		DateTime().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}

func TestFieldOption_Build_ShouldAddCompareOperators_WhenTypeIsDateTimeAndArrayIsDefined(t *testing.T) {
	// Arrange
	opt := NewFieldOption().
		DateTime().
		Array().
		Name("Some Name").
		Label("Some Label")

	// Act
	result, err := opt.Build(make(map[string][]models.Lookup))

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Contains(t, result.Operators, constants.OperatorGreaterThan.String())
	assert.Contains(t, result.Operators, constants.OperatorGreaterThanOrEqual.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThan.String())
	assert.Contains(t, result.Operators, constants.OperatorLessThanOrEqual.String())
}
