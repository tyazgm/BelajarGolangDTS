package helper

import (
	"fmt"
	"sync"
)

func RoutineAcak(wg *sync.WaitGroup, mu *sync.Mutex, i int, data []string, channel chan string) {
	defer wg.Done()
	channel <- fmt.Sprintf("%v %d", data, i)
}

func RoutineRapih(wg *sync.WaitGroup, mu *sync.RWMutex, i int, data []string, channel *chan string) {
	defer wg.Done()
	*channel <- fmt.Sprintf("%v %d", data, i)
}

func Join(wg *sync.WaitGroup, mu *sync.RWMutex, ch1 *chan string, ch2 *chan string, channelJoin *chan string) {
	defer wg.Done()
	mu.Lock()
	*channelJoin <- <-*ch1
	*channelJoin <- <-*ch2
	mu.Unlock()
}
