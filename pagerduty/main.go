package main

import (
	"context"
	"fmt"
	"time"
)

func isCancel(ctx context.Context) {
	timer := time.After(3 * time.Second)
	select {
	case <-timer:
		fmt.Println("task completed")
	case <-ctx.Done():
		fmt.Println("sorry context is cancelled")
	}
}

func isCancel2(ctx context.Context) {
	timer := time.After(5 * time.Second)
	select {
	case <-timer:
		fmt.Println("task completed from second")
	case <-ctx.Done():
		fmt.Println("sorry context is cancelled from second")
	}
}

func main() {
	fmt.Println("hey")
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel context.CancelFunc) {
		time.Sleep(2 * time.Second)
		cancel()
	}(cancel)
	go isCancel(ctx)
	go isCancel2(ctx)

	fmt.Scanln()

}
