package strings_filter

import (
	"strconv"
	"strings"
)

func AddHexPrefix(hex string) string {
	if len(hex) >= 2 && hex[0:2] == "0x" {
		return hex
	}
	return "0x" + hex
}

func Trim0xPrefix(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	}
	return input
}

// ParseIntFromHexStr parse hex string value to int
func ParseIntFromHexStr(value string) (int, error) {
	i, err := strconv.ParseInt(Trim0xPrefix(value), 16, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}
