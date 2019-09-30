package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	numOfErrs     = 5
	numOfFunc     = 40
	numOfRoutines = 30
)

type tickT int

func (t *tickT) testTick(i int) error {
	ticker := time.NewTicker(time.Duration(i) * time.Second)
	fmt.Println(i, "ticktime")
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(100)
			if i%2 != 0 {
				return errors.New("err cause odd number")
			} else {
				return nil
			}
		}
	}
}

func gogo(sl []func() error, toGo, numErrs int) {
	if toGo >= len(sl) {
		return
	}
	errCh := make(chan int, numErrs)
	errSlice := make([]int, 0, numErrs)
	start := make(chan bool)
	for i := 0; i <= toGo; i++ {
		time.Sleep(10 * time.Millisecond)
		go func(i int) {
			<-start
			fmt.Println("goroutine", i)
			for {
				if sl[i]() != nil {
					if len(errSlice) >= numErrs {
						close(errCh)
						return
					}
					errCh <- i
				}
			}
		}(i)
	}
	close(start)
Loop:
	for {
		select {
		case <-errCh:
			errSlice = append(errSlice, <-errCh)
			fmt.Println(len(errSlice), "error length")
			if len(errSlice) >= numErrs {
				break Loop
			}
		}
	}

	fmt.Println("func end")
}

func main() {
	fmt.Println("start")
	sl := make([]func() error, 0)
	for i := 1; i < numOfFunc; i++ {
		time.Sleep(10 * time.Millisecond)
		foo := func(i int) func() error {
			return func() error {
				f := new(tickT)
				return f.testTick(i)
			}
		}(i)
		sl = append(sl, foo)
	}

	gogo(sl, numOfRoutines, numOfErrs)
	fmt.Println("end main")
}
