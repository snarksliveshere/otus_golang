package unpackstr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckStrCorrect(t *testing.T) {
	require.Equal(t, checkStrCorrect("3olala"), false, "Must be an error if the first symbol is num")
	require.Equal(t, checkStrCorrect("olala"), true, "Must be OK")
	require.Equal(t, checkStrCorrect("ASDlala"), true, "Must be OK")
	require.Equal(t, checkStrCorrect("45"), false, "Must be an error if string can be convert to num")
	require.Equal(t, checkStrCorrect("0"), false, "Must be an error if string equal to zero")
	require.Equal(t, checkStrCorrect(""), false, "Must be an error if the string is empty")
}

func TestGetUnpackString(t *testing.T) {
	if s, _ := GetUnpackString("olala3"); s != "olalaaa" {
		t.Fatalf("bad something %v", s)
	}
	if s, _ := GetUnpackString("qwe\\4\\5"); s != "qwe45" {
		t.Fatalf("bad something %v", s)
	}
	if s, _ := GetUnpackString("qwe\\45"); s != "qwe44444" {
		t.Fatalf("bad something %v", s)
	}
	if s, _ := GetUnpackString("q0e5"); s != "eeeee" {
		t.Fatalf("bad something %v", s)
	}
	if s, _ := GetUnpackString("q1e5"); s != "qeeeee" {
		t.Fatalf("bad something %v", s)
	}
	//if s, _ := GetUnpackString("qwe\\\\5"); s != "qwe\\\\\\\\\\" {
	//	t.Fatalf("bad something %q", s)
	//}
}
