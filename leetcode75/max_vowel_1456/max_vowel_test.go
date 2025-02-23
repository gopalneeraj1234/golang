package maxvowel1456

import "testing"

var vowelTests = []struct {
	s    string
	k    int
	want int
}{
	{"abciiidef", 3, 3},
	{"aeiou", 2, 2},
}

func TestMaxVowels(t *testing.T) {
	for _, tt := range vowelTests {
		t.Run(tt.s, func(t *testing.T) {
			got := maxVowels(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("Got %d, but want %d", got, tt.want)
			}
		})
	}
}
