package main

import (
	"fmt"
	"session5/helper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var data1 = []string{"coba1", "coba2", "coba3"}
	var data2 = []string{"bisa1", "bisa2", "bisa3"}

	var ch = make(chan string)

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go helper.RoutineAcak(&wg, &mu, i+1, data1, ch)
		go helper.RoutineAcak(&wg, &mu, i+1, data2, ch)
	}

	for i := 0; i < 8; i++ {
		fmt.Println(<-ch)
	}

	wg.Wait()
}
