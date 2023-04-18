package helper

import (
	"fmt"
	"sync"
)

func RoutineAcak(wg *sync.WaitGroup, mu *sync.Mutex, i int, data []string, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("%v %d", data, i)
}

func RoutineRapih(wg *sync.WaitGroup, mu *sync.RWMutex, i int, data []string, ch2 *chan string) {
	defer wg.Done()
	*ch2 <- fmt.Sprintf("%v %d", data, i)
}

func Join(wg *sync.WaitGroup, mu *sync.RWMutex, ch1 *chan string, ch2 *chan string, chJoin *chan string) {
	defer wg.Done()
	mu.Lock()
	*chJoin <- <-*ch1
	*chJoin <- <-*ch2
	mu.Unlock()
}
