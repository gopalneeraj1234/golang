package greatestcandies1431

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := slices.Max(candies)
	retArr := make([]bool, len(candies))
	for i, candy := range candies {
		if (candy + extraCandies) >= maxValue {
			retArr[i] = true
		}
	}
	return retArr
}
