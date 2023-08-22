package iris_operators

import (
	"errors"
	"fmt"
	"strconv"
)

func compareFloat64Float64(operation string, x, y float64) (bool, error) {
	switch operation {
	case OperatorEq:
		return x == y, nil
	case OperatorNeq:
		return x != y, nil
	case OperatorLt:
		return x < y, nil
	case OperatorLte:
		return x <= y, nil
	case OperatorGte:
		return x >= y, nil
	case OperatorGt:
		return x > y, nil
	}
	return false, errors.New(fmt.Sprintf("invalid operation: %s", operation))
}

func ConvertToFloat64(value any) (float64, bool) {
	switch v := value.(type) {
	case int:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case uint:
		return float64(v), true
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case string:
		newV, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, false
		}
		return newV, true
	}
	return 0, false
}
