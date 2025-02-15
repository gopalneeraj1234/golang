package canflower605

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	got := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestCanPlaceFlowersCantPot(t *testing.T) {
	got := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
	want := false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestCanPlaceFlowersCantPotNext(t *testing.T) {
	got := canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2)
	want := false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
