package stringsclose1657

import (
	"maps"
	"reflect"
	"slices"
	"sort"
)

func createWordMap(word string) map[rune]int {
	wordMap := map[rune]int{}
	for _, ch := range word {
		wordMap[ch]++
	}
	return wordMap
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Map := createWordMap(word1)
	word2Map := createWordMap(word2)

	if !maps.EqualFunc(word1Map, word2Map, func(v1 int, v2 int) bool {
		return true // Compare the keys only
	}) {
		return false
	}

	word1SortedVals := slices.Collect(maps.Values(word1Map))
	sort.Ints(word1SortedVals)
	word2SortedVals := slices.Collect(maps.Values(word2Map))
	sort.Ints(word2SortedVals)
	return reflect.DeepEqual(word1SortedVals, word2SortedVals)
}
