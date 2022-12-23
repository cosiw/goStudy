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

	for i := 0; i < n; i++ {
		go sumation(&sum, i, n)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println(sum)
	fmt.Println(elapsed)

}

func sumation(sumTotal *int64, i int, n int) {

	var sum int64 = 0

	for j := 10000000000 / n * i; j <= 10000000000/n*(i+1); j++ {
		sum += int64(j)
	}
	mutex.Lock()
	*sumTotal += sum
	mutex.Unlock()
	wg.Done()

}
