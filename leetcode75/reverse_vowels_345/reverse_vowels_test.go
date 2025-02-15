package reversevowels345

import "testing"

func TestReverseVowels(t *testing.T) {

	got := reverseVowels("IceCreAm")
	want := "AceCreIm"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseVowelsOneMore(t *testing.T) {

	got := reverseVowels("leetcode")
	want := "leotcede"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseVowelsSingle(t *testing.T) {

	got := reverseVowels("l")
	want := "l"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseVowelsAll(t *testing.T) {

	got := reverseVowels("aeiou")
	want := "uoiea"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
