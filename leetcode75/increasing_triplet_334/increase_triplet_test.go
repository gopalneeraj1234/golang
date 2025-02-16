package increasingtriplet334

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	got := increasingTriplet([]int{1, 2, 3, 4, 5})
	want := true

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}

func TestIncreasingTripletDecrease(t *testing.T) {
	got := increasingTriplet([]int{5, 4, 3, 2, 1})
	want := false

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}
