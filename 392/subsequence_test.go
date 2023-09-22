package subsequence

import "testing"

type substrTest struct {
	Substr   string
	Str      string
	Expected bool
}

var medianTestData = []substrTest{
	{"acf", "abcdef", true},
	{"axc", "ahbgdc", false},
	{"", "abcdef", true},
	{"abc", "", false},
	{"aaaaaa", "bbaaaa", false},
}

func TestIsSubsequence(t *testing.T) {
	for _, test := range medianTestData {
		if result := IsSubsequence(test.Substr, test.Str); result != test.Expected {
			t.Errorf("Expected %v for \"%s\", \"%s\", but got %v", test.Expected, test.Substr, test.Str, result)
		}
	}
}

func BenchmarkIsSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSubsequence(medianTestData[1].Substr, medianTestData[1].Str)
	}
}
