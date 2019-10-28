package drivers

import "fmt"

type InMemStorage struct {
	conn string
}

func (handler *InMemStorage) Execute(str string) (i interface{}) {
	fmt.Println(str)
	return i
}

func NewStorageHandler() *InMemStorage {
	inMemStorage := new(InMemStorage)
	inMemStorage.conn = "connection"
	return inMemStorage
}
