package mb_json_schemas

import (
	"strconv"
	"strings"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func GetStringEvalComparisonResult(operator string, actual, expected string) bool {
	switch operator {
	case "contains":
		if strings.Contains(actual, expected) {
			return true
		}
	case "has-prefix":
		if strings.HasPrefix(actual, expected) {
			return true
		}
	case "has-suffix":
		if strings.HasSuffix(actual, expected) {
			return true
		}
	case "does-not-start-with-any":
		fs := &strings_filter.FilterOpts{
			DoesNotStartWithThese: strings.Split(expected, ","),
		}
		return strings_filter.FilterStringWithOpts(actual, fs)
	case "does-not-include":
		fs := &strings_filter.FilterOpts{
			DoesNotInclude: strings.Split(expected, ","),
		}
		return strings_filter.FilterStringWithOpts(actual, fs)
	case "equals":
		return actual == expected
	case "length-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		actualLen := len(actual)
		return actualLen == expectedLen
	case "length-less-than":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		actualLen := len(actual)
		if actualLen < expectedLen {
			return true
		}
	case "length-less-than-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) <= expectedLen {
			return true
		}
	case "length-greater-than":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) > expectedLen {
			return true
		}
	case "length-greater-than-eq":
		expectedLen := len(expected)
		comparedLengthLimit, err := strconv.Atoi(expected)
		if err == nil {
			expectedLen = comparedLengthLimit
		}
		if len(actual) >= expectedLen {
			return true
		}
	}

	return false
}
