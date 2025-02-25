package equalrow2352

import (
	"strconv"
	"testing"
)

var equalTests = []struct {
	grid [][]int
	want int
}{
	{
		[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1,
	},
	{
		[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3,
	},
}

func TestEqualPairs(t *testing.T) {
	for i, tt := range equalTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := equalPairs(tt.grid)
			if got != tt.want {
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		})
	}

}
