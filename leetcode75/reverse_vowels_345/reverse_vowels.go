package reversevowels345

// Set not available in golang :(
type Set struct {
	elements map[byte]struct{}
}

func (s *Set) contains(c byte) bool {
	_, found := s.elements[c]
	return found
}

func createSet(slice []byte) *Set {
	set := &Set{
		elements: make(map[byte]struct{}),
	}
	for _, element := range slice {
		set.elements[element] = struct{}{}
	}
	return set
}

func getVowelSlice() []byte {
	return []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
}

func reverseVowels(s string) string {
	vowelSet := createSet(getVowelSlice())
	forwardPtr := 0
	backPtr := len(s) - 1

	retArr := make([]byte, len(s))
	for forwardPtr <= backPtr {
		if !vowelSet.contains(s[forwardPtr]) {
			retArr[forwardPtr] = s[forwardPtr]
			forwardPtr++
		} else if !vowelSet.contains(s[backPtr]) {
			retArr[backPtr] = s[backPtr]
			backPtr--
		} else {
			retArr[forwardPtr] = s[backPtr]
			retArr[backPtr] = s[forwardPtr]
			forwardPtr++
			backPtr--
		}
	}
	return string(retArr)
}
