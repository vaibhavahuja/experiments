package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	// Create a context with a cancellation function
	ctx, cancel := context.WithCancel(context.Background())

	// Run a goroutine that cancels the context after 2 seconds
	go func() {
		fmt.Println("starting some work")
		time.Sleep(5 * time.Second) // Simulating some work
		cancel()                    // Cancel the context after 2 seconds
	}()

	for i := 0; ctx.Err() == nil && i < 100; i++ {
		fmt.Println("doing some work")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("context error ", ctx.Err().Error())
	fmt.Println("reached here finally")

}

func doWork(ctx context.Context) {
	// This will return when the context is canceled
	select {
	case <-ctx.Done():
		// If the context is canceled, return
		fmt.Println("Context canceled. Exiting...")
		os.Exit(1)
	default:
	}
}
