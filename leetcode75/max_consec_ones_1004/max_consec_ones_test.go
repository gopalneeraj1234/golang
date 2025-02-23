package maxconsecones1004

import (
	"strconv"
	"testing"
)

var longestOnesTest = []struct {
	nums []int
	k    int
	want int
}{
	{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
	{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
}

func TestLongestOnes(t *testing.T) {

	for i, tt := range longestOnesTest {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := longestOnes(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}
}
