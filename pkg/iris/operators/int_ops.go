package iris_operators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func compareIntInt(operation string, x, y int) (bool, error) {
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

func ConvertToInt(value any) (int, bool) {
	switch v := value.(type) {
	case int:
		return v, true
	case int8:
		return int(v), true
	case int16:
		return int(v), true
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	case uint:
		return int(v), true
	case uint8:
		return int(v), true
	case uint16:
		return int(v), true
	case uint32:
		return int(v), true
	case uint64:
		return int(v), true
	case float32:
		return int(v), true
	case float64:
		return int(v), true
	case string:
		if strings.HasPrefix(v, "0x") {
			dec, err := strings_filter.ParseIntFromHexStr(v)
			if err != nil {
				return 0, false
			}
			return dec, true
		}
		newV, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return newV, true
	}
	return 0, false
}
