package recentcall933

import "testing"

func TestRecentCounter(t *testing.T) {
	obj := Constructor()
	got := obj.Ping(1)
	if got != 1 {
		t.Errorf("Got %d, want %d", got, 1)
	}
	obj.Ping(100)
	obj.Ping(3001)
	got = obj.Ping(3002)
	if got != 3 {
		t.Errorf("Got %d, want %d", got, 3)
	}
}
