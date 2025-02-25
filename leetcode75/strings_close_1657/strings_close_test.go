package stringsclose1657

import "testing"

var strClose = []struct {
	word1 string
	word2 string
	want  bool
}{
	{"abc", "bca", true},
	{"a", "aa", false},
	{"cabbba", "abbccc", true},
	{"uau", "ssx", false},
}

func TestCloseStrings(t *testing.T) {
	for _, tt := range strClose {
		t.Run("Test"+tt.word1+tt.word2, func(t *testing.T) {
			got := closeStrings(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("Got %t, want %t", got, tt.want)
			}
		})
	}
}
