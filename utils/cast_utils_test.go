package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsArray_ShouldReturnFalse_WhenInputIsNotArray(t *testing.T) {
	// Act
	// Arrange
	assert.False(t, IsArray("test"))
	assert.False(t, IsArray(100))
	assert.False(t, IsArray(true))
	assert.False(t, IsArray(time.Now()))
	assert.False(t, IsArray(60))
	assert.False(t, IsArray(struct{}{}))
	assert.False(t, IsArray(make(map[string]interface{})))
}

func TestIsArray_ShouldReturnTrue_WhenInputIsArray(t *testing.T) {
	// Act
	// Arrange
	assert.True(t, IsArray([]string{}))
	assert.True(t, IsArray(make([]string, 0)))
	assert.True(t, IsArray([]float64{}))
	assert.True(t, IsArray(make([]float64, 0)))
	assert.True(t, IsArray([]bool{}))
	assert.True(t, IsArray(make([]bool, 0)))
	assert.True(t, IsArray([]time.Time{}))
	assert.True(t, IsArray(make([]time.Time, 0)))
	assert.True(t, IsArray([]time.Duration{}))
	assert.True(t, IsArray(make([]time.Duration, 0)))
	assert.True(t, IsArray([]string{}))
	assert.True(t, IsArray(make([]string, 0)))
}

func TestArray_ShouldReturnError_WhenInputIsNotValid(t *testing.T) {
	// Arrange
	var data interface{}

	data = "test"

	// Act
	result, err := Array(data)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestArray_ShouldReturnArray_WhenInputIsValid(t *testing.T) {
	// Arrange
	var data interface{}

	data = []string{
		"test1",
		"test2",
	}

	// Act
	result, err := Array(data)

	// Assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "test1", result[0])
	assert.Equal(t, "test2", result[1])
}

func TestIsString_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
	}

	for _, input := range sampleMap {
		// Act
		result := IsString(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsString_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		100,
		int8(10),
		int16(100),
		int32(100),
		int64(100),
		uint(100),
		uint8(10),
		uint16(100),
		uint32(100),
		uint64(100),
		float32(100),
		float64(100),
		true,
		"test",
		time.Now(),
		60,
		nil,
	}

	for input := range sampleMap {
		// Act
		result := IsString(input)

		// Assert
		assert.True(t, result)
	}
}

func TestString_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
	}

	for _, input := range sampleMap {
		// Act
		result, err := String(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestString_ShouldReturnValueAsString_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	now := time.Now()
	sampleMap := map[interface{}]string{
		100:          "100",
		int8(10):     "10",
		int16(100):   "100",
		int32(100):   "100",
		int64(100):   "100",
		uint(100):    "100",
		uint8(10):    "10",
		uint16(100):  "100",
		uint32(100):  "100",
		uint64(100):  "100",
		float32(100): "100",
		float64(100): "100",
		true:         "true",
		"test":       "test",
		now:          now.Format(time.RFC3339),
		60:           "60",
		nil:          "",
	}

	for input, output := range sampleMap {
		// Act
		result, err := String(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, result)
	}
}

func TestIsNumber_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
	}

	for _, input := range sampleMap {
		// Act
		result := IsNumber(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsNumber_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		100,
		int8(10),
		int16(100),
		int32(100),
		int64(100),
		uint(100),
		uint8(10),
		uint16(100),
		uint32(100),
		uint64(100),
		float32(100),
		float64(100),
		true,
		false,
		"123",
		nil,
	}

	for _, input := range sampleMap {
		// Act
		result := IsNumber(input)

		// Assert
		assert.True(t, result)
	}
}

func TestNumber_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
	}

	for _, input := range sampleMap {
		// Act
		result, err := Number(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestNumber_ShouldReturnValueAsNumber_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := map[interface{}]float64{
		100:          100,
		int8(10):     10,
		int16(100):   100,
		int32(100):   100,
		int64(100):   100,
		uint(100):    100,
		uint8(10):    10,
		uint16(100):  100,
		uint32(100):  100,
		uint64(100):  100,
		float32(100): 100,
		float64(100): 100,
		true:         1,
		false:        0,
		"123":        123,
		nil:          0,
	}

	for input, output := range sampleMap {
		// Act
		result, err := Number(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, result)
	}
}

func TestIsBool_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
		60,
		123,
	}

	for _, input := range sampleMap {
		// Act
		result := IsBoolean(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsBool_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		true,
		false,
		"true",
		"false",
		"True",
		"False",
		"TRUE",
		"FALSE",
		1,
		0,
	}

	for _, input := range sampleMap {
		// Act
		result := IsBoolean(input)

		// Assert
		assert.True(t, result)
	}
}

func TestBool_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
		60,
		123,
	}

	for _, input := range sampleMap {
		// Act
		result, err := Boolean(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestBool_ShouldReturnValueAsBoolean_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := map[interface{}]bool{
		true:    true,
		false:   false,
		"true":  true,
		"false": false,
		"True":  true,
		"False": false,
		"TRUE":  true,
		"FALSE": false,
		1:       true,
		0:       false,
	}

	for input, output := range sampleMap {
		// Act
		result, err := Boolean(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, result)
	}
}

func TestIsDate_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		60,
		123,
		true,
		false,
		"12:00:00",
	}

	for _, input := range sampleMap {
		// Act
		result := IsDate(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsDate_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 01, 10, 12, 32, 800, time.UTC),
		"2020-01-01",
		"2020-01-01 10:12:14",
		"2020-01-01 10:12:14.899",
		"2020-01-01T00:00:00Z",
	}

	for _, input := range sampleMap {
		// Act
		result := IsDate(input)

		// Assert
		assert.True(t, result)
	}
}

func TestDate_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		60,
		123,
		true,
		false,
	}

	for _, input := range sampleMap {
		// Act
		result, err := Date(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestDate_ShouldReturnValueAsDate_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := map[interface{}]time.Time{
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC):      time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 01, 10, 12, 32, 800, time.UTC): time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		"2020-01-01":              time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		"2020-01-01 10:12:14":     time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		"2020-01-01 10:12:14.899": time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		"2020-01-01T00:00:00Z":    time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
	}

	for input, output := range sampleMap {
		// Act
		result, err := Date(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, *result)
	}
}

func TestIsTime_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
	}

	for _, input := range sampleMap {
		// Act
		result := IsTime(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsTime_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		60,
		"10:12:11",
		"1H",
		"1H30M",
		"1h30m",
	}

	for _, input := range sampleMap {
		// Act
		result := IsTime(input)

		// Assert
		assert.True(t, result)
	}
}

func TestTime_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		time.Now(),
	}

	for _, input := range sampleMap {
		// Act
		result, err := Time(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestTime_ShouldReturnValueAsTime_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := map[interface{}]int{
		10:         10,
		"1h30m":    1*60*60 + 30*60,
		"01:30:00": 1*60*60 + 30*60,
		"01:30":    1*60*60 + 30*60,
	}

	for input, output := range sampleMap {
		// Act
		result, err := Time(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, *result)
	}
}

func TestIsDateTime_ShouldReturnFalse_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		60,
		123,
		true,
		false,
		"12:00:00",
	}

	for _, input := range sampleMap {
		// Act
		result := IsDateTime(input)

		// Assert
		assert.False(t, result)
	}
}

func TestIsDateTime_ShouldReturnTrue_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 01, 10, 12, 32, 800, time.UTC),
		"2020-01-01",
		"2020-01-01 10:12:14",
		"2020-01-01 10:12:14.899",
		"2020-01-01T00:00:00Z",
	}

	for _, input := range sampleMap {
		// Act
		result := IsDateTime(input)

		// Assert
		assert.True(t, result)
	}
}

func TestDateTime_ShouldReturnError_WhenInputTypeIsNotSupported(t *testing.T) {
	// Arrange
	sampleMap := []interface{}{
		struct{}{},
		"TEST",
		60,
		123,
		true,
		false,
		"12:00:00",
	}

	for _, input := range sampleMap {
		// Act
		result, err := DateTime(input)

		// Assert
		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestDateTime_ShouldReturnValueAsDateTime_WhenInputTypeIsSupported(t *testing.T) {
	// Arrange
	sampleMap := map[interface{}]time.Time{
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC):      time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 01, 10, 12, 32, 800, time.UTC): time.Date(2020, 01, 01, 10, 12, 32, 0, time.UTC),
		"2020-01-01":              time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		"2020-01-01 10:12:14":     time.Date(2020, 01, 01, 10, 12, 14, 0, time.UTC),
		"2020-01-01 10:12:14.899": time.Date(2020, 01, 01, 10, 12, 14, 899000000, time.UTC),
		"2020-01-01T00:00:00Z":    time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
	}

	for input, output := range sampleMap {
		// Act
		result, err := DateTime(input)

		// Assert
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.Equal(t, output, *result)
	}
}
