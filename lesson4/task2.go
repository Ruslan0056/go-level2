package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	signalChanel := make(chan os.Signal)
	go func(ctx context.Context) {
		signal.Notify(signalChanel,
			syscall.SIGTERM)
	}(ctx)

	select {
	case <-signalChanel:
		fmt.Println("Signal terminte triggered.")

 	default:
		fmt.Println("Unknown signal.")
	}

}
