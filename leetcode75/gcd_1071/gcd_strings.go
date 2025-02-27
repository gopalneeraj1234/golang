package gcd_1071

// conditional operator not supported in golang
func getMinStr(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	}
	return str2
}
func gcdOfStrings(str1 string, str2 string) string {
	minStr := getMinStr(str1, str2)
	for checkLen := len(minStr); checkLen > 0; checkLen-- {
		if isRepeated(str1, minStr, checkLen) && isRepeated(str2, minStr, checkLen) {
			return minStr[0:checkLen]
		}
	}
	return ""
}

func isRepeated(str string, minStr string, checkLen int) bool {
	if len(str)%checkLen != 0 {
		return false
	}
	firstStr := minStr[0:checkLen]

	for i := 0; i+checkLen <= len(str); i += checkLen {
		if str[i:i+checkLen] != firstStr {
			return false
		}
	}

	if checkLen <= len(str) {
		return str[0:checkLen] == firstStr
	}
	return true
}
