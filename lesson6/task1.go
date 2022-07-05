package main

import (
	"fmt"
	"sync"
)

const count = 10

func main() {
	var (
		lock sync.Mutex
		ch   = make(chan int)
		done = make(chan struct{})
		wg   = sync.WaitGroup{}
	)

	wg.Add(count)
	for i := 0; i <= count; i += 1 {
		go func() {
			defer close(done)
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			ch <- i
		}()
		sec := <-ch
		fmt.Println(sec)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-done:
				return
			}
		}
	}()

	wg.Wait()

}
