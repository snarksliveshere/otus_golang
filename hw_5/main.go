package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type tickT int

func (t *tickT) testTick(i int) error {
	ticker := time.NewTicker(time.Duration(i) * time.Second)
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().Unix())
			i := rand.Intn(100)
			fmt.Println("inn")
			if i%2 != 0 {
				return errors.New("err cause odd")
			} else {
				return nil
			}
		}
	}
}

func gogo(sl []func() error, toGo, numErrs int) {
	fmt.Println(len(sl), toGo)
	if toGo >= len(sl) {
		return
	}

	errCh := make(chan error)
	for i := 0; i <= toGo; i++ {
		go func() {
			err := sl[i]()
			if err != nil {
				errCh <- err
			}
		}()
	}

	go func() {
		for x := range errCh {
			fmt.Println(x.Error())
		}
	}()

}

func main() {
	fmt.Println("start")
	//ff := func() error {
	//	f1 := new(tickT)
	//	return f1.testTick(2)
	//}
	//var tsl []func() error
	//tsl = append(tsl, ff)
	//tsl[0]()
	//errCh := make(chan error)
	//test1()
	//t := new(tickT)
	//err2 := t.testTick(2)
	sl := make([]func() error, 0)
	for i := 1; i < 5; i++ {
		foo := func() error {
			f := new(tickT)
			return f.testTick(i)
		}
		sl = append(sl, foo)
	}
	err := sl[0]()
	if err != nil {
		fmt.Println("ola err")
	} else {
		fmt.Println("ola ok")

	}

	gogo(sl, 2, 3)
	//fmt.Println(sl)
	//for _, v := range sl {
	//	go func() {
	//		fmt.Println("olala")
	//		err := v()
	//		if err != nil {
	//			errCh <- err
	//		}
	//	}()
	//}
	//
	//go func() {
	//	for x := range errCh {
	//		fmt.Println(x.Error())
	//	}
	//
	//}()

	//if err2 != nil {
	//	fmt.Println("its error testtick")
	//} else {
	//	fmt.Println("its ok testtick")
	//}
	//err := test2()
	//if err != nil {
	//	fmt.Println("its error")
	//} else {
	//	fmt.Println("its ok")
	//}
	//
	//ticker := time.NewTicker(2 * time.Second)
	//
	//ch := make(chan string, 1)
	//
	//go func() {
	//	//time.Sleep(4 * time.Second)
	//	ch <- "Hello"
	//	close(ch)
	//}()
	//timer := time.NewTimer(3 * time.Second)
	//select {
	//case data := <-ch:
	//	fmt.Printf("received %v", data)
	//case <-timer.C:
	//
	//}
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Println("something")
	//		rand.Seed(time.Now().Unix())
	//		i := rand.Intn(100)
	//		fmt.Println("failed to receive in  4s", i, i%2)
	//		//default:
	//		//	fmt.Println("default ticker")
	//	}
	//	//time.Sleep(5 * time.Second)
	//}
	//
	//test()
}

func test() (err error) {
	var m map[string]string
	fmt.Println("test")
	if _, ok := m["ola"]; !ok {
		fmt.Println("err")
		return err
	}
	return nil
}

func test2() error {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			rand.Seed(time.Now().Unix())
			i := rand.Intn(100)
			fmt.Println(i)
			if i%2 != 0 {
				return errors.New("err cause odd")
			} else {
				return nil
			}
		}
	}
}

func test1() {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		time.Sleep(4 * time.Second)
		for i := 40; i <= 50; i++ {
			ch <- i
			if i == 43 {
				close(ch)
				break
			}

		}
	}()

	go func() {
		for x := range ch {
			fmt.Println(x)
		}

	}()

}
