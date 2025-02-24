package findpivot724

import (
	"strconv"
	"testing"
)

var pivotTests = []struct {
	nums []int
	want int
}{
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{1, 2, 3}, -1},
	{[]int{2, 1, -1}, 0},
}

func TestPivotIndex(t *testing.T) {
	for i, tt := range pivotTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Errorf("Got %d, Want %d", got, tt.want)
			}
		})
	}
}
