// Package test
// @author tabuyos
// @since 2023/7/21
// @description test
package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCurrying(t *testing.T) {
	fmt.Println("hello")
}

func TestRoutine(t *testing.T) {
	ss := make(map[int]sync.Mutex)
	ss[0] = sync.Mutex{}
	ss[1] = sync.Mutex{}
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		mu = ss[(i%2)+1]

		mu.Lock()

		go func(counter int) {
			fmt.Println(counter)
			time.Sleep(1e9)
			mu.Unlock()
		}(i)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%v\n", input)
	}
}

func TestChannel0(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("... %d\n", i)
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Printf("1> %v\n", <-ch)
		}
	}()

	go func() {
		for {
			fmt.Printf("2> %v\n", <-ch)
		}
	}()

	time.Sleep(1e9)
}

func TestChannel1(t *testing.T) {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println(111)
		ch <- i
	}
}
