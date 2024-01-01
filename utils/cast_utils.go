package utils

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/filtex/filtex-go/errors"
)

func IsArray(val interface{}) bool {
	kind := reflect.TypeOf(val).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

func Array(value interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0)

	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(value)

		for j := 0; j < s.Len(); j++ {
			result = append(result, s.Index(j).Interface())
		}
	default:
		return nil, errors.NewCouldNotBeCastedError()
	}

	return result, nil
}

func IsString(val interface{}) bool {
	_, err := String(val)
	return err == nil
}

func String(val interface{}) (string, error) {
	switch v := val.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(v), nil
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case time.Time:
		return v.Format(time.RFC3339), nil
	case *time.Time:
		return v.Format(time.RFC3339), nil
	case time.Duration:
		return v.String(), nil
	case *time.Duration:
		return v.String(), nil
	case reflect.Value:
		if v.CanInterface() {
			return String(v.Interface())
		}

		switch v.Kind() {
		case reflect.String:
			return v.String(), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return String(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return String(v.Uint())
		case reflect.Float32, reflect.Float64:
			return String(v.Float())
		case reflect.Bool:
			return String(v.Bool())
		}

		return "", errors.NewCouldNotBeCastedError()
	case nil:
		return "", nil
	default:
		return "", errors.NewCouldNotBeCastedError()
	}
}

func IsNumber(val interface{}) bool {
	_, err := Number(val)
	return err == nil
}

func Number(val interface{}) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v == true {
			return float64(1), nil
		}
		return float64(0), nil
	case string:
		return strconv.ParseFloat(val.(string), 64)
	case reflect.Value:
		if v.CanInterface() {
			return Number(v.Interface())
		}

		switch v.Kind() {
		case reflect.String:
			return Number(v.String())
		case reflect.Bool:
			return Number(v.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(v.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return float64(v.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return v.Float(), nil
		}

		return 0.0, errors.NewCouldNotBeCastedError()
	case nil:
		return 0.0, nil
	default:
		return 0.0, errors.NewCouldNotBeCastedError()
	}
}

func IsBoolean(val interface{}) bool {
	_, err := Boolean(val)
	return err == nil
}

func Boolean(value interface{}) (bool, error) {
	if value == nil {
		return false, nil
	}

	s, err := String(value)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}

	return b, nil
}

func IsDate(val interface{}) bool {
	_, err := Date(val)
	return err == nil
}

func Date(val interface{}) (*time.Time, error) {
	datetime, err := DateTime(val)
	if err != nil {
		return nil, err
	}

	date := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())

	return &date, nil
}

func IsTime(val interface{}) bool {
	_, err := Time(val)
	return err == nil
}

func Time(val interface{}) (*int, error) {
	if val == nil {
		return nil, errors.NewCouldNotBeCastedError()
	}

	str, err := String(val)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(strings.ToLower(str))
	if err == nil {
		seconds := int(duration.Seconds())
		return &seconds, nil
	}

	timeWithSeconds, err := time.Parse("15:04:05", str)
	if err == nil {
		zeroTime, _ := time.Parse("15:04:05", "00:00:00")
		seconds := int(timeWithSeconds.Sub(zeroTime).Seconds())
		return &seconds, nil
	}

	timeWithoutSeconds, err := time.Parse("15:04", str)
	if err == nil {
		zeroTime, _ := time.Parse("15:04", "00:00")
		seconds := int(timeWithoutSeconds.Sub(zeroTime).Seconds())
		return &seconds, nil
	}

	sec, err := strconv.Atoi(str)
	if err == nil {
		return &sec, nil
	}

	return nil, errors.NewCouldNotBeCastedError()
}

func IsDateTime(val interface{}) bool {
	_, err := DateTime(val)
	return err == nil
}

func DateTime(val interface{}) (*time.Time, error) {
	if val == nil {
		return nil, nil
	}

	var formats = []string{
		"2006-01-02 15:04:05 Z0700 MST",
		"2006-01-02 15:04:05 Z07:00 MST",
		"2006-01-02 15:04:05 Z0700 -0700",
		"Mon Jan _2 15:04:05 -0700 MST 2006",
		time.RFC822Z, // "02 Jan 06 15:04 -0700"
		time.RFC3339, // "2006-01-02T15:04:05Z07:00", RFC3339Nano
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05 Z07:00",
		"2006-01-02 15:04:05.999999999 -0700 MST",
		time.RubyDate, // "Mon Jan 02 15:04:05 -0700 2006"
		time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC822,   // "02 Jan 06 15:04 MST",
		"2006-01-02 15:04:05 MST",
		time.UnixDate, // "Mon Jan _2 15:04:05 MST 2006",
		time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST",
		time.RFC850,   // "Monday, 02-Jan-06 15:04:05 MST",
		time.Kitchen,  // "3:04PM"
		"01/02/06 15:04",
		time.Stamp, // "Jan _2 15:04:05", time.StampMilli, time.StampMicro, time.StampNano,
		time.ANSIC, // "Mon Jan _2 15:04:05 2006"
		"2006-01-02 15:04",
		"2006-01-02T15:04",
		"01/02/2006 15:04",
		"01/02/06 15:04:05",
		"01/02/2006 15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"_2/Jan/2006 15:04:05",
		"2006-01-02",
	}

	s, err := String(val)
	if err != nil {
		return nil, err
	}

	for _, format := range formats {
		t, err := time.Parse(format, s)
		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.NewCouldNotBeCastedError()
}
