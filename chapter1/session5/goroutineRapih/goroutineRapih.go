package main

import (
	"fmt"
	"session5/helper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	var data1 = []string{"coba1", "coba2", "coba3"}
	var data2 = []string{"bisa1", "bisa2", "bisa3"}

	var chanJoin = make(chan string, 2)
	var chan1 = make(chan string)
	var chan2 = make(chan string)

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go helper.RoutineRapih(&wg, &mu, i+1, data1, &chan1)
		go helper.RoutineRapih(&wg, &mu, i+1, data2, &chan2)
	}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go helper.Join(&wg, &mu, &chan1, &chan2, &chanJoin)
	}

	for i := 0; i < 8; i++ {
		fmt.Println(<-chanJoin)
	}

	go func() {
		wg.Wait()
		close(chanJoin)
	}()
}
