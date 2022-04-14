package g

import (
	"errors"
	"strconv"
)

func ForceInt(s string, defaultInt int) int {
	parsed, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultInt
	}
	return int(parsed)
}

func ForceInt32(s string, defaultInt32 int32) int32 {
	parsed, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultInt32
	}
	return int32(parsed)
}

func ForceInt64(s string, defaultInt64 int64) int64 {
	parsed, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultInt64
	}
	return int64(parsed)
}

// String converts the input value to a string.
// You can pass with a precision number (defaults to 12) if
// you are converting the float to string.
//   String(3.1415926, 10)
//   String(3)
func String(args ...any) string {
	value := args[0]
	var precision int = 12 // default

	switch value.(type) {
	case string:
		v, _ := value.(string)
		return v
	case int:
		v, _ := value.(int)
		return strconv.Itoa(v)
	case int32:
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10)
	case int64:
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10)
	case float32:
		v, _ := value.(float32)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(float64(v), 'f', precision, 64)
	case float64:
		v, _ := value.(float64)
		if len(args) >= 2 {
			precision = args[1].(int)
		}
		return strconv.FormatFloat(v, 'f', precision, 64)
	default:
		return ""
	}
}

func Int(value any) (int, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.Atoi(v)
	case int:
		v, _ := value.(int)
		return v, nil
	case int32:
		v, _ := value.(int32)
		return int(v), nil
	case int64:
		v, _ := value.(int64)
		return int(v), nil
	case float32:
		v, _ := value.(float32)
		return int(v), nil
	case float64:
		v, _ := value.(float64)
		return int(v), nil
	default:
		return int(0), errors.New("Unknown type")
	}
}

func Int32(value any) (int32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseInt(v, 10, 32)
		return int32(result), err
	case int:
		v, _ := value.(int)
		return int32(v), nil
	case int32:
		v, _ := value.(int32)
		return int32(v), nil
	case int64:
		v, _ := value.(int64)
		return int32(v), nil
	case float32:
		v, _ := value.(float32)
		return int32(v), nil
	case float64:
		v, _ := value.(float64)
		return int32(v), nil
	default:
		return int32(0), errors.New("Unknown type")
	}
}

func Int64(value any) (int64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseInt(v, 10, 64)
	case int:
		v, _ := value.(int)
		return int64(v), nil
	case int32:
		v, _ := value.(int32)
		return int64(v), nil
	case int64:
		v, _ := value.(int64)
		return v, nil
	case float32:
		v, _ := value.(float32)
		return int64(v), nil
	case float64:
		v, _ := value.(float64)
		return int64(v), nil
	default:
		return int64(0), errors.New("Unknown type")
	}
}

func Float32(value any) (float32, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseFloat(v, 32)
		return float32(result), err
	case int:
		v, _ := value.(int)
		return float32(v), nil
	case int32:
		v, _ := value.(int32)
		return float32(v), nil
	case int64:
		v, _ := value.(int64)
		return float32(v), nil
	case float32:
		v, _ := value.(float32)
		return v, nil
	case float64:
		v, _ := value.(float64)
		return float32(v), nil
	default:
		return float32(0), errors.New("Unknown type")
	}
}

func Float64(value any) (float64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	default:
		return float64(0), errors.New("Unknown type")
	}
}
