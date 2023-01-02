package strings_filter

func AddHexPrefix(hex string) string {
	if len(hex) >= 2 && hex[0:2] == "0x" {
		return hex
	}
	return "0x" + hex
}
