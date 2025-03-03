package wordcounter

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	t.Log("Testing with files in the testcases directory")
	wordCounts := map[string]*CountResult{
		"singleword.txt":   {4, 1, 1},
		"doubleword.txt":   {9, 1, 2},
		"onlynewlines.txt": {5, 6, 0},
		"list.txt":         {23, 5, 5},
		"startspaces.txt":  {13, 1, 1},
	}

	for testfilename := range wordCounts {
		wordCounter := NewWordCounter("testcases/" + testfilename)
		countResult, err := wordCounter.Count()
		if err != nil {
			t.Fatalf("Fatal error %q", err)
		}
		if !reflect.DeepEqual(countResult, wordCounts[testfilename]) {
			t.Errorf("NewWordCounter(%q)=%q, want %q", testfilename, countResult, wordCounts[testfilename])
		}
	}
}
