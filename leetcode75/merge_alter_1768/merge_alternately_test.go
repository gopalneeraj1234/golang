package mergealter1768

import "testing"

func TestMergeAlternatelyEqStr(t *testing.T) {
	got := mergeAlternately("test", "test")
	want := "tteesstt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMergeAlternatelyNonEq(t *testing.T) {
	got := mergeAlternately("abc", "abcde")
	want := "aabbccde"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMergeAlternatelyZeroLen(t *testing.T) {
	got := mergeAlternately("", "")
	want := ""

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
