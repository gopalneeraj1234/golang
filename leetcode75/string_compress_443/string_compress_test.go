package stringcompress443

import (
	"fmt"
	"testing"
)

func checkSliceEquals(got []byte, want []byte, t *testing.T) {
	for i, element := range want {
		if element != got[i] {
			t.Errorf("got %d at %d but want %d", got[i], i, element)
		}
	}
}

func TestCompress(t *testing.T) {
	sentSlice := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(sentSlice)
	got := compress(sentSlice)
	wantSlicePart := []byte{'a', '2', 'b', '2', 'c', '3'}
	want := 6

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
	fmt.Println(sentSlice)
	checkSliceEquals(sentSlice, wantSlicePart, t)

}

func TestCompressSingle(t *testing.T) {
	sentSlice := []byte{'a'}
	fmt.Println(sentSlice)
	got := compress(sentSlice)
	wantSlicePart := []byte{'a'}
	want := 1

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
	fmt.Println(sentSlice)
	checkSliceEquals(sentSlice, wantSlicePart, t)

}

func TestCompressDoubleDigit(t *testing.T) {
	sentSlice := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(sentSlice)
	got := compress(sentSlice)
	wantSlicePart := []byte{'a', 'b', '1', '2'}
	want := 4

	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
	fmt.Println(sentSlice)
	checkSliceEquals(sentSlice, wantSlicePart, t)

}
