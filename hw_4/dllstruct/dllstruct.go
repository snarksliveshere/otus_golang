package dllstruct

type List struct {
	items []*item
	ptr   int
}

type item struct {
	Value interface{}
}

func (l *List) Len() int {
	return len(l.items)
}

func (l *List) GetItems() []*item {
	return l.items
}

func (l *List) SetItems(items []*item) {
	l.items = items
}

func (l *List) GetItemByIndex(index int) *item {
	if l.isZeroLen() {
		return nil
	}
	if index >= l.Len() {
		return nil
	}
	return l.GetItems()[index]
}

func (l *List) setPtr(index int) {
	l.ptr = index
}
func (l *List) getPtr() int {
	return l.ptr
}

func (l *List) First() *item {
	l.setPtr(0)
	return l.GetItemByIndex(0)
}

func (l *List) Next() *item {
	n := l.getPtr() + 1
	l.setPtr(n)
	if item := l.GetItemByIndex(n); item != nil {
		return item
	}
	l.setPtr(n - 1)
	return nil
}

func (l *List) Prev() *item {
	if l.getPtr() == 0 {
		return nil
	}
	n := l.getPtr() - 1
	l.setPtr(n)
	if item := l.GetItemByIndex(n); item != nil {
		return item
	}
	l.setPtr(n + 1)
	return nil
}

func (l *List) Last() *item {
	l.setPtr(l.Len() - 1)
	return l.GetItemByIndex(l.Len() - 1)
}

func (l *List) PushBack(value interface{}) int {
	item := item{value}
	l.items = append(l.items, &item)
	return l.Len()
}

func (l *List) PushFront(value interface{}) int {
	i := append([]*item(nil), &item{value})
	l.items = append(i, l.items...)
	return l.Len()
}

func (l *List) RemoveFront() *item {
	if l.isZeroLen() {
		return nil
	}
	value := l.items[0:1][0]
	l.items = l.items[1:]
	if l.isZeroLen() {
		return nil
	}
	return value
}

func (l *List) Remove(i int) *item {
	if item := l.GetItemByIndex(i); item != nil {
		l.items = append(l.items[0:i], l.items[i+1:]...)
		return item
	}
	return nil
}

func (l *List) RemoveBack() *item {
	if l.isZeroLen() {
		return nil
	}
	value := l.items[l.Len()-1]
	l.items = l.items[:l.Len()-1]
	if l.isZeroLen() {
		return nil
	}
	return value
}

func (l *List) isZeroLen() bool {
	if l.Len() == 0 {
		return true
	}
	return false
}
