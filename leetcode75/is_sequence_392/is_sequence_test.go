package issequence392

import (
	"fmt"
	"math"
	"testing"

	"math/rand"
)

func TestIsSubSequence(t *testing.T) {

	got := isSubsequence("abc", "ahbgdc")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsSubSequenceNot(t *testing.T) {

	got := isSubsequence("abcd", "ahbgdc")
	want := false

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsSubSequenceSingle(t *testing.T) {

	got := isSubsequence("a", "a")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsSubSequenceLong(t *testing.T) {

	longStr := make([]rune, int(math.Pow(10, 4)))
	for range int(math.Pow(10, 4)) {
		longStr = append(longStr, rune('a'+rand.Intn(25)))
	}

	searchStr := make([]rune, 100)

	for range 100 {
		searchStr = append(searchStr, rune('a'+rand.Intn(25)))
	}

	fmt.Println(string(searchStr))
	fmt.Println(string(longStr))
	got := isSubsequence(string(searchStr), string(longStr))
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
