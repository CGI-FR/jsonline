package jsonline

import (
	"fmt"
	"strconv"
)

const (
	conversionSize = 64
)

func toString(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case []byte:
		return string(typedValue)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func toNumeric(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int64, int, int16, int8, byte, rune, float64, float32:
		return typedValue
	case bool:
		if typedValue {
			return 1
		}

		return 0
	case string:
		r, err := strconv.ParseFloat(typedValue, conversionSize)
		if err == nil {
			return r
		}

		return fmt.Sprintf("ERROR: %v", err)
	case []byte:
		return toNumeric(fmt.Sprintf("%v", string(typedValue)))
	default:
		return toNumeric(fmt.Sprintf("%v", val))
	}
}

func toBoolean(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case int64, int, int16, int8, byte, rune:
		return typedValue != 0
	case float64, float32:
		return typedValue != 0.0
	case bool:
		return typedValue
	case string:
		r, err := strconv.ParseBool(typedValue)
		if err == nil {
			return r
		}

		f64, err := strconv.ParseFloat(typedValue, conversionSize)
		if err == nil {
			return toBoolean(f64)
		}

		return fmt.Sprintf("ERROR: %v", err)
	case []byte:
		return toBoolean(fmt.Sprintf("%v", string(typedValue)))
	default:
		return toBoolean(fmt.Sprintf("%v", val))
	}
}
