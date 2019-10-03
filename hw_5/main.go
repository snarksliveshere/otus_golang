package main

import (
	"errors"
	"fmt"
	"runtime"
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

func go2(sl []func() error, numOfRoutines, numOfErrors int) {
	tasks := make(chan func() error, numOfFunc)
	errCh := make(chan int, numOfErrors)
	die := make(chan bool)
	for _, v := range sl {
		tasks <- v
	}
	time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	for i := 0; i <= numOfRoutines; i++ {
		go func(i int) {
			wg.Add(1)
			worker(tasks, errCh, die, &wg)
		}(i)
	}

	wg.Wait()
	fmt.Println("ola")
	for {
		select {

		case <-errCh:
		default:
			die <- true
			//case <-time.After(time.Millisecond * 100):
			//	fmt.Println(len(errCh), "len errch", time.Now())
			//	fmt.Printf("Runtime numGoroutine %d\n", runtime.NumGoroutine())
			// 1 горутина - это main, следовательно, все остальные отвалились
			//if runtime.NumGoroutine() == 1 {
			//	return
			//}
		}
	}
}

func worker(tasks <-chan func() error, errCh chan int, die <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working")
	for {
		select {
		case <-die:
			return
		case <-tasks:
			f := <-tasks
			if f() != nil {
				fmt.Println(len(errCh), "err length")
				if _, ok := <-die; ok {
					return
				}
				errCh <- 1
			}
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

	//gogo(sl, numOfRoutines, numOfErrs)
	//fmt.Println("end main")
}
