package containermostwater11

import (
	"math"
	"math/rand"
	"testing"
)

func TestMaxArea(t *testing.T) {
	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	want := 49

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestMaxAreaSmall(t *testing.T) {
	got := maxArea([]int{1, 1})
	want := 1

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestMaxAreaLarge(t *testing.T) {
	height := make([]int, int(math.Pow(10, 4)))
	for i := 0; i < int(math.Pow(10, 4)); i++ {
		height[i] = rand.Intn(int(math.Pow(10, 4)))
	}
	maxArea(height)
}
