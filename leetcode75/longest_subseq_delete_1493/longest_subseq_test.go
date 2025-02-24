package longestsubseqdelete1493

import (
	"strconv"
	"testing"
)

var longsubtests = []struct {
	nums []int
	want int
}{
	{[]int{1, 1, 0, 1}, 3},
	{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
	{[]int{1, 1, 1}, 2},
}

func TestLongestSubarray(t *testing.T) {

	for i, tt := range longsubtests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := longestSubarray(tt.nums)
			if got != tt.want {
				t.Errorf("Got %d, Want %d", got, tt.want)
			}
		})
	}
}
