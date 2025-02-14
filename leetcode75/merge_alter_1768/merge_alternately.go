package mergealter1768

import "strings"

func appendRest(sb *strings.Builder, word string, indexToStart int) {
	for i := indexToStart; i < len(word); i++ {
		sb.WriteByte(word[i])
	}
}

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder

	var minLength = min(len(word1), len(word2))
	for i := 0; i < minLength; i++ {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}

	appendRest(&sb, word1, minLength)
	appendRest(&sb, word2, minLength)

	return sb.String()
}
