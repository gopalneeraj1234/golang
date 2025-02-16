package increasingtriplet334

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32

	for _, n := range nums { // Without _ it will give the index not the number
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}
