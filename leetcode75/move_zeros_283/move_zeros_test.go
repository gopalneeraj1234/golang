package movezeros283

import "testing"

func checkSliceEquals(got []int, want []int, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("got %d len, want %d", len(got), len(want))
	}
	for i, element := range want {
		if element != got[i] {
			t.Errorf("got %d at %d but want %d", got[i], i, element)
		}
	}
}

func TestMoveZeros(t *testing.T) {
	sent := []int{0, 1, 0, 3, 12}
	moveZeroes(sent)
	want := []int{1, 3, 12, 0, 0}
	checkSliceEquals(sent, want, t)
}

func TestMoveZerosSingle(t *testing.T) {
	sent := []int{0}
	moveZeroes(sent)
	want := []int{0}
	checkSliceEquals(sent, want, t)
}

func TestMoveZerosSorted(t *testing.T) {
	sent := []int{0, 0, 0, 3, 12}
	moveZeroes(sent)
	want := []int{3, 12, 0, 0, 0}
	checkSliceEquals(sent, want, t)
}

func TestMoveZerosNothing(t *testing.T) {
	sent := []int{3, 12, 0}
	moveZeroes(sent)
	want := []int{3, 12, 0}
	checkSliceEquals(sent, want, t)
}

func TestMoveZerosAll(t *testing.T) {
	sent := []int{0, 0, 0}
	moveZeroes(sent)
	want := []int{0, 0, 0}
	checkSliceEquals(sent, want, t)
}
