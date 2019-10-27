package infrastructure

import (
	"fmt"
)

type InMemStorage struct {
	Conn string
}

func (handler *InMemStorage) Execute(interface{}) (i interface{}) {
	fmt.Println("execute")
	return i
}

func (handler *InMemStorage) ExecuteBulk(interface{}) (i []interface{}) {
	fmt.Println("execute bulk")
	return i
}

func NewStorageHandler() *InMemStorage {
	inMemStorage := new(InMemStorage)
	inMemStorage.Conn = "connection"
	return inMemStorage
}
