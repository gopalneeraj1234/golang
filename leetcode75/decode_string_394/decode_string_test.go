package decodestring394

import "testing"

var decodeTests = []struct {
	s    string
	want string
}{
	{"3[a]2[bc]", "aaabcbc"},
	{"3[a2[c]]", "accaccacc"},
	{"2[abc]3[cd]ef", "abcabccdcdcdef"},
}

func TestDecodeString(t *testing.T) {
	for _, tt := range decodeTests {
		t.Run("Test"+tt.s, func(t *testing.T) {
			got := decodeString(tt.s)
			if got != tt.want {
				t.Errorf("Got %q, Want %q", got, tt.want)
			}
		})
	}
}
