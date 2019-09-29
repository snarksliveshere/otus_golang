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
			fmt.Println(i)
			if i%2 != 0 {
				return errors.New("err cause odd")
			} else {
				return nil
			}
		}
	}
}

func main() {
	fmt.Println("olala")
	//test1()
	t := new(tickT)
	err2 := t.testTick(2)
	if err2 != nil {
		fmt.Println("its error testtick")
	} else {
		fmt.Println("its ok testtick")
	}
	err := test2()
	if err != nil {
		fmt.Println("its error")
	} else {
		fmt.Println("its ok")
	}

	ticker := time.NewTicker(2 * time.Second)

	ch := make(chan string, 1)

	go func() {
		//time.Sleep(4 * time.Second)
		ch <- "Hello"
		close(ch)
	}()
	timer := time.NewTimer(3 * time.Second)
	select {
	case data := <-ch:
		fmt.Printf("received %v", data)
	case <-timer.C:

	}
	for {
		select {
		case <-ticker.C:
			fmt.Println("something")
			rand.Seed(time.Now().Unix())
			i := rand.Intn(100)
			fmt.Println("failed to receive in  4s", i, i%2)
			//default:
			//	fmt.Println("default ticker")
		}
		//time.Sleep(5 * time.Second)
	}

	test()
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
