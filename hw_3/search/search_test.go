package search

import (
	"testing"
)

var casesNormalizeStr = []struct {
	in  string
	out string
}{
	{"\t\tola", " ola"},
	{"         ", " "},
	{"ABC", "abc"},
	{",ola", "ola"},
	{".a", "a"},
	{"a\nb", "ab"},
}

func TestNormalizeStr(t *testing.T) {
	for _, c := range casesNormalizeStr {
		out := normalizeStr(c.in)
		if out != c.out {
			t.Errorf("NormalizeStr(%q) == %q, want %q", c.in, out, c.out)
		}
	}
}

func TestGetWordsByNum(t *testing.T) {
	cases := []struct {
		num  int
		word []word
		out  []string
	}{
		{1, []word{{"first", 1}, {"second", 2}}, []string{"first"}},
		{10, []word{{"first", 1}, {"second", 2}}, []string{"first", "second"}},
		{2, []word{{"first", 1}, {"second", 2}, {"zero", 0}}, []string{"first", "second"}},
	}

	for _, c := range cases {
		out := getWordsByNum(c.num, c.word)
		if !compareStringSlices(out, c.out) {
			t.Errorf("getWordsByNum() == %q, want %q", out, c.out)
		}
	}
}

func compareStringSlices(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
