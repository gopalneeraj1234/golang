package issequence392

func isSubsequence(s string, t string) bool {
	sPtr := 0
	tPtr := 0

	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			sPtr++
			tPtr++
		} else {
			tPtr++
		}
	}
	return sPtr == len(s)
}
