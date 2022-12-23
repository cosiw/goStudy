package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, a int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick", a)
		case <-terminate:
			fmt.Println("Terminated!", a)
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square : %d, %d \n", n*n, a)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go square(&wg, ch, 1)
	go square(&wg, ch, 2)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
