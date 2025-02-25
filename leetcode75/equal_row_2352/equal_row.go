package equalrow2352

import (
	"strconv"
	"strings"
)

func getStringFromArray(arr []int, sep string) string {
	if len(arr) == 0 {
		return ""
	}

	strSlice := make([]string, len(arr))
	for i, num := range arr {
		strSlice[i] = strconv.Itoa(num)
	}
	return strings.Join(strSlice, sep)
}

func equalPairs(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])
	rowMap := make(map[string]int)
	for _, rowSlice := range grid {
		rowMap[getStringFromArray(rowSlice, ",")]++
	}

	numMatches := 0
	for col := 0; col < colLen; col++ {
		colSlice := make([]int, rowLen)
		for row := 0; row < rowLen; row++ {
			colSlice[row] = grid[row][col]
		}
		rowMatches, found := rowMap[getStringFromArray(colSlice, ",")]
		if found {
			numMatches += rowMatches
		}
	}

	return numMatches
}
