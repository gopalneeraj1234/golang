package productarray238

import "testing"

func checkSliceEquals(got []int, want []int, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("got %d, want %d", len(got), len(want))
	}
	for i, element := range got {
		if element != want[i] {
			t.Errorf("got %d at %d but want %d", element, i, want[i])
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	got := productExceptSelf([]int{1, 2, 3, 4})
	want := []int{24, 12, 8, 6}

	checkSliceEquals(got, want, t)
}

func TestProductExceptSelfZeros(t *testing.T) {
	got := productExceptSelf([]int{-1, 1, 0, -3, 3})
	want := []int{0, 0, 9, 0, 0}

	checkSliceEquals(got, want, t)
}
