package dllstruct

import (
	"testing"
)

func TestLen(t *testing.T) {
	cases := []struct {
		in  []*item
		out int
	}{
		{[]*item{{1}}, 1},
		{[]*item{{1}, {item{2}}}, 2},
		{[]*item{}, 0},
	}
	for _, c := range cases {
		var l = List{}
		l.SetItems(c.in)
		out := l.Len()
		if out != c.out {
			t.Errorf("len() == %q, want %q", out, c.out)
		}
	}
}

func TestPushBack(t *testing.T) {
	cases := []struct {
		in  []*item
		val interface{}
		out int
	}{
		{in: []*item{}, val: 1, out: 1},
		{in: []*item{{1}}, val: 2, out: 2},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.PushBack(c.val)
		if out != c.out {
			t.Errorf("PushBack() == %q, want %q", out, c.out)
		}
	}
}

func TestPushFront(t *testing.T) {
	cases := []struct {
		in  []*item
		val interface{}
		out int
	}{
		{in: []*item{}, val: 1, out: 1},
		{in: []*item{{1}}, val: 2, out: 2},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.PushFront(c.val)
		if out != c.out {
			t.Errorf("PushFront() == %q, want %q", out, c.out)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		in  []*item
		val int
		out interface{}
	}{
		{in: []*item{}, val: 1, out: nil},
		{in: []*item{{2}}, val: 1, out: 2},
		{in: []*item{{"front"}, {"back"}}, val: 1, out: "back"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.Remove(c.val)
		if out != nil {
			if out.Value != c.out {
				t.Errorf("Remove() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestRemoveFront(t *testing.T) {
	cases := []struct {
		in  []*item
		out interface{}
	}{
		{in: []*item{}, out: nil},
		{in: []*item{{1}}, out: nil},
		{in: []*item{{"front"}, {"back"}}, out: "front"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.RemoveFront()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("RemoveFront() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestRemoveBack(t *testing.T) {
	cases := []struct {
		in  []*item
		out interface{}
	}{
		{in: []*item{}, out: nil},
		{in: []*item{{1}}, out: nil},
		{in: []*item{{"front"}, {"back"}}, out: "back"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.RemoveBack()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("RemoveBack() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestFirst(t *testing.T) {
	cases := []struct {
		in  []*item
		out interface{}
	}{
		{in: []*item{}, out: nil},
		{in: []*item{{1}}, out: 1},
		{in: []*item{{"front"}, {"back"}}, out: "front"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.First()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("First() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		in  []*item
		out interface{}
	}{
		{in: []*item{}, out: nil},
		{in: []*item{{1}}, out: 1},
		{in: []*item{{"front"}, {"back"}}, out: "back"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		out := l.Last()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("Last() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestNext(t *testing.T) {
	cases := []struct {
		in  []*item
		ptr int
		out interface{}
	}{
		{in: []*item{{1}, {2}}, ptr: 0, out: 2},
		{in: []*item{{1}, {2}}, ptr: 1, out: nil},
		{in: []*item{{"front"}, {"back"}, {"next"}}, ptr: 1, out: "next"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		l.setPtr(c.ptr)
		out := l.Next()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("Next() == %q, want %q", out, c.out)
			}
		}
	}
}

func TestPrev(t *testing.T) {
	cases := []struct {
		in  []*item
		ptr int
		out interface{}
	}{
		{in: []*item{{1}, {2}}, ptr: 1, out: 1},
		{in: []*item{{1}, {2}}, ptr: 0, out: nil},
		{in: []*item{{"front"}, {"back"}, {"next"}}, ptr: 2, out: "back"},
	}
	for _, c := range cases {
		l := setItems(c.in)
		l.setPtr(c.ptr)
		out := l.Prev()
		if out != nil {
			if out.Value != c.out {
				t.Errorf("Prev() == %q, want %q", out, c.out)
			}
		}
	}
}

func setItems(items []*item) List {
	l := List{}
	l.SetItems(items)
	return l
}
