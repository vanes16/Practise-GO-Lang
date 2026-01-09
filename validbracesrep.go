package validbracesrep

import "strings"

func ValidBraces(str string) bool {
	for {
		newStr := strings.ReplaceAll(str, "()", "")
		newStr = strings.ReplaceAll(newStr, "[]", "")
		newStr = strings.ReplaceAll(newStr, "{}", "")

		// tidak ada perubahan → berhenti
		if newStr == str {
			break
		}

		str = newStr
	}

	return str == ""
}
