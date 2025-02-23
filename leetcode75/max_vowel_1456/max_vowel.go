package maxvowel1456

func maxVowels(s string, k int) int {
	windowStart := 0
	windowEnd := k - 1

	numVowels := getNumVowels(windowStart, windowEnd, s)
	retMax := numVowels
	windowStart++
	windowEnd++
	for windowEnd < len(s) {
		if isVowel(s[windowStart-1]) {
			numVowels--
		}
		if isVowel(s[windowEnd]) {
			numVowels++
		}
		retMax = max(numVowels, retMax)
		windowEnd++
		windowStart++
	}
	return retMax
}

func getNumVowels(windowStart, windowEnd int, s string) int {
	numVowels := 0
	for i := windowStart; i <= windowEnd; i++ {
		if isVowel(s[i]) {
			numVowels++
		}
	}
	return numVowels
}

var vowelSet = map[byte]struct{}{
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

func isVowel(b byte) bool {
	_, found := vowelSet[b]
	return found
}
