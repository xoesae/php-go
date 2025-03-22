package utils

import (
	"fmt"
	"strconv"
)

func ToFloat(value interface{}) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case string:
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			return num
		}
	}
	panic(fmt.Sprintf("Conversion error: %v", value))
}

func ToBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v != ""
	}
	return false
}

func IsNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	default:
		return false
	}
}
