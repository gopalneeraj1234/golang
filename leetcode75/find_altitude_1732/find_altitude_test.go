package findaltitude1732

import (
	"strconv"
	"testing"
)

var altitudesTestData = []struct {
	gain []int
	want int
}{
	{[]int{-5, 1, 5, 0, -7}, 1},
	{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
}

func TestLargestAltitude(t *testing.T) {
	for i, tt := range altitudesTestData {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := largestAltitude(tt.gain)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}
}
