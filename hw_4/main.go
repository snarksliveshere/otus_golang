package main

import (
	"fmt"
	l "github.com/otus_golang/hw_4/dllstruct"
)

func main() {
	list := l.List{}
	list.PushBack("second")
	list.PushBack("third")
	list.PushFront("first")

	fmt.Println(list.Len())                   // 3
	fmt.Println(list.GetItemByIndex(0).Value) // first
	fmt.Println(list.Remove(12))              // first

	fmt.Println(list.First().Value) // first
	fmt.Println(list.Prev())        // nil , ptr to first
	fmt.Println(list.Next().Value)  // second

	fmt.Println(list.Last().Value) // third
	fmt.Println(list.Next())       // nil, ptr to last
	fmt.Println(list.Prev().Value) // second

	fmt.Println(list.RemoveFront().Value) // first
	fmt.Println(list.RemoveBack().Value)  // third

	fmt.Println(list.GetItems())
	fmt.Println(list.Remove(0).Value)
	fmt.Println(list.GetItems())

}
