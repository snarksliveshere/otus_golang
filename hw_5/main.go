package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	numOfErrs     = 5
	numOfFunc     = 10
	numOfRoutines = 2
)

type tickT int

// эта конструкция просто для удобства
func (t *tickT) testTimer(i int) error {
	timer := time.NewTimer(time.Duration(i) * time.Second)
	fmt.Println(i, " int func")
	for {
		select {
		case <-timer.C:
			//return nil
			return errors.New("just err")
		}
	}
}

func go2(sl []func() error, numOfRoutines, numOfErrors int) {
	lenSl := len(sl)
	if numOfRoutines >= lenSl {
		numOfRoutines = lenSl
	}
	tasks := make(chan func() error, lenSl)
	errCh := make(chan int, numOfErrors)
	for _, v := range sl {
		tasks <- v
	}
	var wg sync.WaitGroup
	for i := 0; i < numOfRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			go worker(tasks, errCh, &wg, numOfErrors, lenSl, i)
		}(i)

	}

	wg.Wait()
}

func worker(tasks <-chan func() error, errCh chan int, wg *sync.WaitGroup, numOfErrs, sliceLength, index int) {
	defer wg.Done()
	die := make(chan bool, sliceLength)

	for {
		// index передаю просто для наглядности
		fmt.Println("start working routine # ", index)
		select {
		case <-die:
			fmt.Println("buy!")
			return
		case f := <-tasks:
			if f() != nil {
				fmt.Println("scratch up errors", len(errCh))
				if len(errCh) == numOfErrs {
					die <- true
					fmt.Println("good buy!")
					return
				}
				errCh <- 1
			}
		default:
			return
		}
	}
}

func main() {
	fmt.Println("start")
	sl := make([]func() error, 0)
	for i := 1; i < numOfFunc; i++ {
		foo := func(i int) func() error {
			return func() error {
				f := new(tickT)
				return f.testTimer(i)
			}
		}(i)
		sl = append(sl, foo)
	}
	go2(sl, numOfRoutines, numOfErrs)

	fmt.Println("end working")
}
