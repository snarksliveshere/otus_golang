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
		err := CopySubStr(c.from, c.to, c.limit, c.offset, "y")
		if err != nil {
			t.Errorf("TestLimitWithOffset(), err while start function")
		}
		f, err := os.Open(c.to)
		if err != nil {
			t.Errorf("TestLimitWithOffset(), err while Open file %s", c.to)
		}
		fs, err := f.Stat()
		if err != nil {
			t.Errorf("TestLimitWithOffset(), err while get stat file %s", c.to)
		}
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
		err := CopySubStr(c.from, c.to, c.limit, c.offset, "y")
		if err != nil {
			t.Errorf("TestLimitZeroOffsetZero(), err while start function")
		}
		f, err := os.Open(c.to)
		if err != nil {
			t.Errorf("TestLimitZeroOffsetZero(), err while Open file %s", c.to)
		}
		fs, err := f.Stat()
		if err != nil {
			t.Errorf("TestLimitZeroOffsetZero(), err while get stat file %s", c.to)
		}
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
		eof           string
	}{
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    100,
			offset:   400,
			destSize: 44,
			eof:      "y",
		},
		{
			from:     "../files/if.txt",
			to:       "../files/of.txt",
			limit:    100,
			offset:   400,
			destSize: 44,
			eof:      "n",
		},
	}

	for _, c := range cases {
		err := CopySubStr(c.from, c.to, c.limit, c.offset, c.eof)
		f, errF := os.Open(c.to)
		if errF != nil {
			t.Errorf("TestLimitWithOffset(), err while Open file %s", c.to)
		}
		fs, errS := f.Stat()
		if errS != nil {
			t.Errorf("TestLimitWithOffset(), err while get stat file %s", c.to)
		}
		if c.eof == "n" && fs.Size() == c.destSize {
			t.Errorf("TestEOF() limit == %d, offset %d", c.limit, c.offset)
		}
		if fs.Size() != c.destSize && err != nil && c.eof == "y" {
			t.Errorf("TestEOF() limit == %d, offset %d", c.limit, c.offset)
		}
		if fs.Size() != c.destSize && err == nil {
			t.Errorf("TestEOF() limit == %d, offset %d", c.limit, c.offset)
		}
	}
}

func TestExplicitErrors(t *testing.T) {
	cases := []struct {
		from, to      string
		limit, offset int64
		res           string
		destSize      int64
		eof           string
	}{
		{
			from:   "../files/if.txt",
			to:     "../files/of.txt",
			limit:  100,
			offset: 500,
			eof:    "y",
		},
		{
			from:   "../files/if.txt",
			to:     "../files/of.txt",
			limit:  100,
			offset: 400,
			eof:    "n",
		},
		{
			from:   "../files/if1.txt",
			to:     "../files/of.txt",
			limit:  10,
			offset: 20,
			eof:    "y",
		},
		{
			from:   "../files/if.txt",
			to:     "$dsfsd../files/of.txt",
			limit:  10,
			offset: 20,
			eof:    "y",
		},
	}

	for _, c := range cases {
		err := CopySubStr(c.from, c.to, c.limit, c.offset, c.eof)
		if err == nil {
			t.Errorf("TestExplicitErrors() limit == %d, offset %d", c.limit, c.offset)
		}
	}
}
