package gocopy

import (
	"os"
	"testing"
)

func TestLimitWithOffset(t *testing.T) {
	cases := []struct {
		from, to      string
		limit, offset int64
		res           string
		destSize      int64
	}{
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    90,
			offset:   100,
			destSize: 90,
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    85,
			offset:   100,
			destSize: 85,
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    10,
			offset:   0,
			destSize: 10,
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    10,
			offset:   10,
			destSize: 10,
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    0,
			offset:   10,
			destSize: 0,
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    149,
			offset:   149,
			destSize: 149,
		},
	}

	for _, c := range cases {
		CopySubStr(c.from, c.to, c.limit, c.offset)
		f, _ := os.Open(c.to)
		fs, _ := f.Stat()
		if fs.Size() != c.destSize {
			t.Errorf("TestLimitWithOffset() limit == %d, offset %d", c.limit, c.offset)
		}
	}
}

func TestLimitZeroOffsetZero(t *testing.T) {
	cases := []struct {
		from, to      string
		limit, offset int64
		res           string
		destSize      int64
	}{
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    0,
			offset:   0,
			destSize: 0,
		},
	}

	for _, c := range cases {
		CopySubStr(c.from, c.to, c.limit, c.offset)
		f, _ := os.Open(c.to)
		fs, _ := f.Stat()
		if fs.Size() != c.destSize {
			t.Errorf("TestLimitZeroOffsetZero() limit == %d, offset %d", c.limit, c.offset)
		}
	}
}

func TestEOF(t *testing.T) {
	cases := []struct {
		from, to      string
		limit, offset int64
		res           string
		destSize      int64
	}{
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    100,
			offset:   400,
			destSize: 44,
		},
	}

	for _, c := range cases {
		CopySubStr(c.from, c.to, c.limit, c.offset)
		f, _ := os.Open(c.to)
		fs, _ := f.Stat()
		if fs.Size() != c.destSize {
			t.Errorf("TestEOF() limit == %d, offset %d", c.limit, c.offset)
		}
	}
}
