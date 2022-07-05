package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

const count = 10

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		counter int64
		lock    sync.Mutex
	)

	go func() {
		for i := 1; ; i++ {
			fmt.Println("Бесконечный цикл", i)
			if i%10 == 0 {
				runtime.Gosched()
				fmt.Println("Бесконечный цикл прерван")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second)

	for i := 0; i < count; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter += 1
			fmt.Println("Горутина", counter)
		}()
	}
	time.Sleep(time.Second)
}
