package greatestcandies1431

import "testing"

func checkSliceEquals(got []bool, want []bool, t *testing.T) {
	if len(got) != len(want) {
		t.Errorf("got %d, want %d", len(got), len(want))
	}
	for i, element := range got {
		if element != want[i] {
			t.Errorf("got %t at %q but want %t", element, i, want[i])
		}
	}
}
func TestKidsWithCandies(t *testing.T) {
	got := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	want := []bool{true, true, true, false, true}
	checkSliceEquals(got, want, t)
}

func TestKidsWithCandiesManyFalse(t *testing.T) {
	got := kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	want := []bool{true, false, false, false, false}
	checkSliceEquals(got, want, t)
}
