package uniqueoccur1207

import (
	"strconv"
	"testing"
)

var uniqueTests = []struct {
	nums []int
	want bool
}{
	{[]int{1, 2, 2, 1, 1, 3}, true},
	{[]int{1, 2}, false},
	{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
}

func TestUniqueOccurrences(t *testing.T) {

	for i, tt := range uniqueTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := uniqueOccurrences(tt.nums)
			if got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}
