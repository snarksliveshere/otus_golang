package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

const (
	numOfErrs     = 5
	numOfFunc     = 10
	numOfRoutines = 8
)

type tickT int

// эта конструкция просто для удобства
func (t *tickT) testTimer(i int) error {
	timer := time.NewTimer(time.Duration(i) * time.Second)
	for {
		select {
		case <-timer.C:
			return errors.New("just err")
		}
	}
}

func gogo(sl []func() error, toGo, numErrs int) {
	if toGo >= len(sl) {
		return
	}
	errCh := make(chan int, numErrs)
	for i := 0; i <= toGo; i++ {
		go func(i int) {
			for {
				if sl[i]() != nil {
					select {
					case errCh <- i:
					default:
						return
					}
				}

			}
		}(i)
	}
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			fmt.Println(len(errCh), "len errch", time.Now())
			fmt.Printf("Runtime numGoroutine %d\n", runtime.NumGoroutine())
			// 1 горутина - это main, следовательно, все остальные отвалились
			if runtime.NumGoroutine() == 1 {
				return
			}
		}
	}
}

func main() {
	fmt.Println("start")
	sl := make([]func() error, 0)
	for i := 1; i < numOfFunc; i++ {
		time.Sleep(10 * time.Millisecond)
		foo := func(i int) func() error {
			return func() error {
				f := new(tickT)
				return f.testTimer(i)
			}
		}(i)
		sl = append(sl, foo)
	}

	gogo(sl, numOfRoutines, numOfErrs)
	fmt.Println("end main")
}
