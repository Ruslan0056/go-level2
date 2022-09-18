package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const count = 10

func main() {
	var (
		counter int64
		wg      = sync.WaitGroup{}
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Горутина", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Количество горутин -", counter)
}
