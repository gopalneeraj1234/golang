package reversewords151

import "strings"

func reverseWords(s string) string {

	splitStrs := strings.Split(s, " ")

	var sb strings.Builder

	for i := len(splitStrs) - 1; i >= 0; i-- {
		if len(splitStrs[i]) == 0 {
			continue
		}
		sb.WriteString(splitStrs[i])
		if i > 0 {
			sb.WriteByte(' ')
		}
	}

	return strings.TrimSpace(sb.String())
}
