package reversewords151

import "testing"

func TestReverseWords(t *testing.T) {

	got := reverseWords("the sky is blue")
	want := "blue is sky the"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseWordsSingleWord(t *testing.T) {

	got := reverseWords("the")
	want := "the"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseWordsSingleLetter(t *testing.T) {

	got := reverseWords("t")
	want := "t"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseWordsSingleLetterMultiple(t *testing.T) {

	got := reverseWords("t s e f")
	want := "f e s t"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReverseWordsSingleLetterSpaces(t *testing.T) {

	got := reverseWords("   t s e f   ")
	want := "f e s t"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
