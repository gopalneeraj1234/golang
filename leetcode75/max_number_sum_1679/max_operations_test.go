package maxnumbersum1679

import "testing"

func TestMaxOperations(t *testing.T) {
	got := maxOperations([]int{1, 2, 3, 4}, 5)
	want := 2

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}

func TestMaxOperationsMore(t *testing.T) {
	got := maxOperations([]int{3, 1, 3, 4, 3}, 6)
	want := 1

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}
