package unpackstr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckStrCorrect(t *testing.T) {
	require.Equal(t, checkStrCorrect("olala"), true, "Must be OK - normal entry string")
	require.Equal(t, checkStrCorrect("ASDlala"), true, "Must be OK normal entry string with capitalized")
	require.Equal(t, checkStrCorrect("ASDlala45"), true, "Must be OK normal entry string with numbers")
	require.Equal(t, checkStrCorrect("3olala"), false, "Must be an error if the first symbol is num")
	require.Equal(t, checkStrCorrect("45"), false, "Must be an error if string can be convert to num")
	require.Equal(t, checkStrCorrect("0"), false, "Must be an error if string equal to zero")
	require.Equal(t, checkStrCorrect(""), false, "Must be an error if the string is empty")
}

func TestSymbolFilter(t *testing.T) {
	require.Equal(t, skipNotSymbol(`\`, `a`, `no`), true, "Just a backslash")
	require.Equal(t, skipNotSymbol(`4`, `\`, `\\`), true, "Must be number after double backslash - quantifier")
	require.Equal(t, skipNotSymbol(`4`, `a`, `no`), true, "Must be Number quantifier")
	require.Equal(t, skipNotSymbol(`4`, `\`, `no`), false, "Must be a symbol - escaping number")
	require.Equal(t, skipNotSymbol(`\`, `\`, `no`), false, "Must be a symbol - escaping backslash")
	require.Equal(t, skipNotSymbol(`a`, `\`, `no`), false, "Must be a symbol - escaping letter")
}

// вот теперь не совсем понятно, как его тестировать
//func TestSymbolDict(t *testing.T) {
//	a := []map[string]interface{
//		{"index": 0, "symbol": "a"},
//		{"index": 1, "symbol": "b"},
//		{"index": 2, "symbol": "c"},
//	}
//	require.Equal(t, symbolDict("abc3"), a, "Just a backslash")
//}

func TestGetPrevSymbol(t *testing.T) {
	str1, str2, str3, str4 := `1\\to`, "abcdef", "abc", "ab"
	i1, i2, i3, i4 := 2, 4, 1, 0
	if prev, slashes := getPrevSymbol(i1, str1); prev != `\` && slashes != `1\` {
		t.Fatalf("an error has occured %s %s %s %s", str1, prev, slashes, string(str1[i1]))
	}
	if prev, slashes := getPrevSymbol(i2, str2); prev != "d" && slashes != "cd" {
		t.Fatalf("an error has occured %s %s %s %s", str2, prev, slashes, string(str2[i2]))
	}
	if prev, slashes := getPrevSymbol(i3, str3); prev != "a" && slashes != "" {
		t.Fatalf("an error has occured %s %s %s %s", str3, prev, slashes, string(str3[i3]))
	}
	if prev, slashes := getPrevSymbol(i4, str4); prev != "" && slashes != "" {
		t.Fatalf("an error has occured %s %s %s %s", str4, prev, slashes, string(str4[i4]))
	}
}

func TestGetUnpackString(t *testing.T) {
	if s, err := GetUnpackString("45"); err == nil {
		t.Fatalf("an error has occured %v", s)
	}
	if s, err := GetUnpackString("4aaa"); err == nil {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString("a4bc2d5e"); s != "aaaabccddddde" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString("abcd"); s != "abcd" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString(`qwe\4\5`); s != "qwe45" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString(`qwe\45`); s != "qwe44444" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString("q0e5"); s != "eeeee" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString("q1e5"); s != "qeeeee" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString(`qwe\\5`); s != `qwe\\\\\` {
		t.Fatalf("an error has occured slashes %v", s)
	}
	if s, _ := GetUnpackString(`qwe\00\5`); s != `qwe5` {
		t.Fatalf("an error has occured 2 %v", s)
	}
	if s, _ := GetUnpackString("q01e5"); s != "qeeeee" {
		t.Fatalf("an error has occured %v", s)
	}
	if s, _ := GetUnpackString(`qwe\01\5`); s != `qwe05` {
		t.Fatalf("an error has occured 1 %v", s)
	}
}
