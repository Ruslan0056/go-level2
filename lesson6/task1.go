package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

const count = 10

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		lock sync.Mutex
		ch   = make(chan int)
		wg = sync.WaitGroup{}
	)

	wg.Add(count)
	for i := 0; i <= count; i += 1 {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
		
			ch <- rand.Intn(100)
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= count; i++ {
			v := <-ch
			fmt.Println(v)

		}

	}()

	wg.Wait()

}
