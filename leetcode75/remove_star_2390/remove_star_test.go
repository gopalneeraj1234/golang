package removestar2390

import "testing"

var starTests = []struct {
	s    string
	want string
}{
	{"leet**cod*e", "lecoe"},
	{"erase*****", ""},
}

func TestRemoveStars(t *testing.T) {
	for _, tt := range starTests {
		t.Run("Test"+tt.s, func(t *testing.T) {
			got := removeStars(tt.s)
			if got != tt.want {
				t.Errorf("Got %q, Want %q", got, tt.want)
			}
		})
	}
}
