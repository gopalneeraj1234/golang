package finddiffarr2215

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var diffArrTests = []struct {
	nums1 []int
	nums2 []int
	want  [][]int
}{
	{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{
		{1, 3},
		{4, 6},
	}},
}

func TestFindDiffArr(t *testing.T) {
	for i, tt := range diffArrTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := findDifference(tt.nums1, tt.nums2)
			for _, arr := range got {
				sort.Ints(arr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("The arrays don't match")
			}
		})
	}
}
