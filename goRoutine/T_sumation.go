package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	var sum int64 = 0
	start := time.Now()
	n := 5
	wg.Add(n)
	for i := 0; i < 10000000000; i += 10000000000 / n {

		go sumation(&sum, i, i+10000000000/n)
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println(sum, elapsed)

}

func sumation(sumTotal *int64, start, end int) {
	var sum int64
	for i := start; i < end; i++ {
		sum += int64(i)
	}

	mutex.Lock()
	*sumTotal += sum
	defer mutex.Unlock()
	wg.Done()

}
