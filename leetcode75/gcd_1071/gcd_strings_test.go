package gcd_1071

import "testing"

func TestGcdOfStringsFull(t *testing.T) {

	got := gcdOfStrings("ABCABC", "ABC")
	want := "ABC"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGcdOfStringsPart(t *testing.T) {

	got := gcdOfStrings("ABABAB", "ABAB")
	want := "AB"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGcdOfStringsInvalid(t *testing.T) {

	got := gcdOfStrings("LEET", "CODE")
	want := ""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGcdOfStringsMatchFirst(t *testing.T) {

	got := gcdOfStrings("ABCDEF", "ABC")
	want := ""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
