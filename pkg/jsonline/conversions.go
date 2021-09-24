package jsonline

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
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

func toBinary(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case []byte:
		return base64.StdEncoding.EncodeToString(typedValue)
	case string:
		return base64.StdEncoding.EncodeToString([]byte(typedValue))
	default:
		return base64.StdEncoding.EncodeToString([]byte(toString(typedValue).(string)))
	}
}

func toDateTime(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case time.Time:
		return typedValue.Format(time.RFC3339)
	case string:
		t, err := time.Parse(time.RFC3339, typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toDateTime(t)
	default:
		return toDateTime(toString(val))
	}
}

func toTime(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case time.Time:
		return typedValue.Format("15:04:05Z07:00")
	case string:
		t, err := time.Parse("15:04:05Z07:00", typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toDateTime(t)
	default:
		return toDateTime(toString(val))
	}
}

func toTimeStamp(val interface{}) interface{} {
	switch typedValue := val.(type) {
	case nil:
		return nil
	case time.Time:
		return typedValue.UnixMilli()
	case int64:
		return typedValue
	case string:
		t, err := time.Parse(time.RFC3339, typedValue)
		if err != nil {
			return fmt.Sprintf("ERROR: %v", err)
		}

		return toDateTime(t)
	default:
		return toDateTime(toString(val))
	}
}
