package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ExchangeRate struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

var wg sync.WaitGroup

func getData(c chan *ExchangeRate) error {
	resp, err := http.Get("https://api.frankfurter.app/latest?from=KRW")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resJson := ExchangeRate{}
	if err := json.Unmarshal(data, &resJson); err != nil {
		return err
	}

	c <- &resJson
	wg.Done()
	return nil
}

func main() {
	c := make(chan *ExchangeRate)
	var sum float64 = 0
	var avg float64
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go getData(c)
		data := <-c

		sum += data.Rates["USD"]
		avg = sum / float64(i)
		fmt.Println(avg)
	}

	wg.Wait()

	// for range time.Tick(time.Second * 1) {
	// 	fmt.Println(data.Rates["USD"])
	// }
}
