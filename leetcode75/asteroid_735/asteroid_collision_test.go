package asteroid735

import (
	"reflect"
	"strconv"
	"testing"
)

var asteroidTests = []struct {
	asteroids []int
	want      []int
}{
	{[]int{5, 10, -5}, []int{5, 10}},
	{[]int{8, -8}, []int{}},
	{[]int{10, 2, -5}, []int{10}},
	{[]int{-2, -2, 1, -2}, []int{-2, -2, -2}},
}

func TestAsteroidCollision(t *testing.T) {
	for i, tt := range asteroidTests {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got len %d, want len %d", len(got), len(tt.want))
			}
		})
	}
}
