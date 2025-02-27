package maxvowel1456

func maxVowels(s string, k int) int {
	windowStart := 0
	windowEnd := k - 1

	vowels := getvowels(windowStart, windowEnd, s)
	retMax := vowels
	windowStart++
	windowEnd++
	for windowEnd < len(s) {
		if isVowel(s[windowStart-1]) {
			vowels--
		}
		if isVowel(s[windowEnd]) {
			vowels++
		}
		retMax = max(vowels, retMax)
		windowEnd++
		windowStart++
	}
	return retMax
}

func getvowels(windowStart, windowEnd int, s string) int {
	vowels := 0
	for i := windowStart; i <= windowEnd; i++ {
		if isVowel(s[i]) {
			vowels++
		}
	}
	return vowels
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
