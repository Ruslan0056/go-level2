package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var sum int64
	pool := make(chan struct{}, runtime.NumCPU())
	for i := 0; i < 1000; i++ {
		pool <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-pool
			}()
			atomic.AddInt64(&sum, 1)

		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
