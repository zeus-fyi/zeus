package strings_filter

import "strings"

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
