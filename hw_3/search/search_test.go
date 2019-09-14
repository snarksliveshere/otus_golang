package search

import (
	"testing"
)

func TestNormalizeStr(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
		out := normalizeStr(c.in)
		if out != c.out {
			t.Errorf("NormalizeStr(%q) == %q, want %q", c.in, out, c.out)
		}
	}
}

func TestGetDicts(t *testing.T) {
	cases := []struct {
		in  string
		out []word
	}{
		{"olala num olala", []word{{"olala", 2}, {"num", 1}}},
		{"first second second", []word{{"first", 1}, {"second", 2}}},
	}
	for _, c := range cases {
		out := getDicts(c.in)
		if !compareWordStruct(out, c.out) {
			t.Errorf("getWordsByNum() == %q, want %q", out, c.out)
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

func compareWordStruct(a, b []word) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].word != b[i].word && a[i].freq != b[i].freq {
			return false
		}
	}

	return true
}
