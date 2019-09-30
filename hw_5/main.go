package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type tickT int

func (t *tickT) testTick() error {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(100)
			if i%2 != 0 {
				return errors.New("err cause odd")
			} else {
				return nil
			}
		}
	}
}

func (t *tickT) testTick2(i int) error {
	ticker := time.NewTicker(time.Duration(i) * time.Second)
	fmt.Println(i, "ticktime")
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(100)
			if i%2 != 0 {
				return errors.New("err cause odd")
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
			if sl[i]() != nil {
				if len(errSlice) >= numErrs {
					close(errCh)
					return
				}
				errCh <- i
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
	for i := 1; i < 40; i++ {
		time.Sleep(10 * time.Millisecond)
		foo := func(i int) func() error {
			return func() error {
				f := new(tickT)
				return f.testTick2(i)
			}
		}(i)
		sl = append(sl, foo)
	}

	gogo(sl, 30, 5)
	fmt.Println("end main")
}
