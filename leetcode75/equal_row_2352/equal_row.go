package equalrow2352

import (
	"strconv"
	"strings"
)

func getStringFromArray(arr []int, sep string) string {
	if len(arr) == 0 {
		return ""
	}

	strs := make([]string, len(arr))
	for i, num := range arr {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, sep)
}

func equalPairs(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])
	rowToCount := make(map[string]int)
	for _, gridRow := range grid {
		rowToCount[getStringFromArray(gridRow, ",")]++
	}

	matches := 0
	for col := range colLen {
		gridColumn := make([]int, rowLen)
		for row := range rowLen {
			gridColumn[row] = grid[row][col]
		}
		rowMatches, found := rowToCount[getStringFromArray(gridColumn, ",")]
		if found {
			matches += rowMatches
		}
	}

	return matches
}
