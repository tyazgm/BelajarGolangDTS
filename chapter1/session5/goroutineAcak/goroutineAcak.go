package main

import (
	"fmt"
	"session5/helper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var dataCoba = []string{"coba1", "coba2", "coba3"}
	var dataBisa = []string{"bisa1", "bisa2", "bisa3"}

	var channel = make(chan string)

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go helper.RoutineAcak(&wg, &mu, i+1, dataCoba, channel)
		go helper.RoutineAcak(&wg, &mu, i+1, dataBisa, channel)
	}

	for i := 0; i < 8; i++ {
		fmt.Println(<-channel)
	}

	wg.Wait()
}
