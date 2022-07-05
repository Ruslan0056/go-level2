package main

import (
	"fmt"
	"sync"
)

const count = 10

func main() {
	var (
		counter int64
		wg      = sync.WaitGroup{}
		lock sync.Mutex
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter += 1
			fmt.Println("Горутина", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Количество горутин -", counter)
}
