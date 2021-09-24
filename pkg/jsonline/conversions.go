package jsonline

import (
	"fmt"
	"strconv"
)

const (
	conversionBase = 10
	conversionSize = 64
)

func toString(val interface{}) interface{} {
	if val == nil {
		return nil
	} else if b, ok := val.([]byte); ok {
		return string(b)
	} else {
		return fmt.Sprintf("%v", val)
	}
}

func toNumeric(val interface{}) interface{} {
	if val == nil {
		return nil
	}

	switch typedValue := val.(type) {
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
