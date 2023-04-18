package main

import (
	"fmt"
	"session5/helper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	var dataCoba = []string{"coba1", "coba2", "coba3"}
	var dataBisa = []string{"bisa1", "bisa2", "bisa3"}

	var chanJoin = make(chan string, 2)
	var chanCoba = make(chan string)
	var chanBisa = make(chan string)

	for i := 0; i < 4; i++ {
		wg.Add(2)
		go helper.RoutineRapih(&wg, &mu, i+1, dataCoba, &chanCoba)
		go helper.RoutineRapih(&wg, &mu, i+1, dataBisa, &chanBisa)
	}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go helper.Join(&wg, &mu, &chanCoba, &chanBisa, &chanJoin)
	}

	for i := 0; i < 8; i++ {
		fmt.Println(<-chanJoin)
	}

	go func() {
		wg.Wait()
		close(chanJoin)
	}()
}
