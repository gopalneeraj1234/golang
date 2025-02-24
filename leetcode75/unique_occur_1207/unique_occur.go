package uniqueoccur1207

func uniqueOccurrences(arr []int) bool {
	elements := map[int]int{}

	for _, num := range arr {
		curr, found := elements[num]
		if found {
			elements[num] = curr + 1
		} else {
			elements[num] = 1
		}
	}

	uniqSet := map[int]struct{}{}
	for _, val := range elements {
		_, found := uniqSet[val]
		if found {
			return false
		}
		uniqSet[val] = struct{}{}
	}
	return true
}
