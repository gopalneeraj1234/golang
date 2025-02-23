package maxaverage643

import "testing"

func TestFindMaxAverage(t *testing.T) {
	got := findMaxAverage([]int{1, 12, -5, -6, 50}, 4)
	want := 12.75

	if got != want {
		t.Errorf("Got %f, want %f", got, want)
	}
}

func TestFindMaxAverageSmall(t *testing.T) {
	got := findMaxAverage([]int{5}, 1)
	want := 5.0

	if got != want {
		t.Errorf("Got %f, want %f", got, want)
	}
}
