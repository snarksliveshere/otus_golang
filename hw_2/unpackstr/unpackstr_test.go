package unpackstr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckStrCorrect(t *testing.T) {
	require.Equal(t, CheckStrCorrect("3olala"), false, "Must be an error if the first symbol is num")
	require.Equal(t, CheckStrCorrect("olala"), true, "Must be OK")
	require.Equal(t, CheckStrCorrect("ASDlala"), true, "Must be OK")
	require.Equal(t, CheckStrCorrect("45"), false, "Must be an error if string can be convert to num")
	require.Equal(t, CheckStrCorrect("0"), false, "Must be an error if string equal to zero")
	require.Equal(t, CheckStrCorrect(""), false, "Must be an error if the string is empty")
}
