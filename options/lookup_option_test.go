package options

import (
	"testing"

	"github.com/filtex/filtex-go/models"
	"github.com/stretchr/testify/assert"
)

func TestNewLookupOption_ShouldReturnLookupOption(t *testing.T) {
	// Act
	opt := NewLookupOption()

	// Assert
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.key)
	assert.NotNil(t, opt.values)
}

func TestLookupOption_Key_ShouldSetKeyAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewLookupOption()

	// Act
	result := opt.Key("some_key")

	// Assert
	assert.NotNil(t, result)
	assert.NotNil(t, result.key)
	assert.NotNil(t, result.values)
	assert.Equal(t, "some_key", result.key)
	assert.Equal(t, opt.values, result.values)
}

func TestLookupOption_Values_ShouldSetValuesAndReturnItself(t *testing.T) {
	// Arrange
	opt := NewLookupOption()

	lookups := []models.Lookup{
		{
			Name:  "Enabled",
			Value: true,
		},
		{
			Name:  "Disabled",
			Value: false,
		},
	}

	// Act
	result := opt.Values(lookups)

	// Assert
	assert.NotNil(t, result)
	assert.NotNil(t, result.key)
	assert.NotNil(t, result.values)
	assert.Equal(t, opt.key, result.key)
	assert.Equal(t, lookups, result.values)
}

func TestLookupOption_Build_ShouldReturnError_WhenKeyIsNotDefined(t *testing.T) {
	// Arrange
	opt := NewLookupOption().
		Values([]models.Lookup{
			{
				Name:  "Enabled",
				Value: true,
			},
			{
				Name:  "Disabled",
				Value: false,
			},
		})

	// Act
	result, err := opt.Build()

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestLookupOption_Build_ShouldReturnError_WhenValuesAreNotDefined(t *testing.T) {
	// Arrange
	opt := NewLookupOption().
		Key("some_key")

	// Act
	result, err := opt.Build()

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestLookupOption_Build_ShouldReturnLookupMap_WhenKeyAndValuesAreDefined(t *testing.T) {
	// Arrange
	opt := NewLookupOption().
		Key("some_key").
		Values([]models.Lookup{
			{
				Name:  "Enabled",
				Value: true,
			},
			{
				Name:  "Disabled",
				Value: false,
			},
		})

	// Act
	result, err := opt.Build()

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.NotNil(t, result["some_key"])
	assert.Len(t, result["some_key"], 2)
}
