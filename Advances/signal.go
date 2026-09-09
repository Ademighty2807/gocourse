package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())

	sigChan := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-sigChan
		fmt.Printf("\nReceived signal: %s\n", sig)
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("Stopping work due to signal.")
			fmt.Println("Graceful exit")
			// os.Exit(0)
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

// func main() {
// 	pid := os.Getpid()
// 	fmt.Printf("Process ID: %d\n", pid)
// 	sigChan := make(chan os.Signal, 1)
// 	done := make(chan bool, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

// 	go func() {
// 		sig := <-sigChan
// 		fmt.Printf("Received signal: %s\n", sig)
// 		done <- true
// 	}()

// 	go func() {

// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("Stopping work due to signal.")
// 				return
// 			default:
// 				// Simulate work
// 				fmt.Println("Working...")
// 				time.Sleep(time.Second)
// 			}
// 			fmt.Println("Graceful exit")
// 			os.Exit(0)

// 		}

// 		// for sigChan := range sigChan {
// 		// 	switch sigChan {
// 		// 	case syscall.SIGINT:
// 		// 		fmt.Println("Received SIGINT (Interrupt)")
// 		// 	case syscall.SIGTERM:
// 		// 		fmt.Println("Received SIGTERM (Terminate)")
// 		// 	case syscall.SIGHUP:
// 		// 		fmt.Println("Received SIGHUP (Hangup)")
// 		// 	}
// 		// 	fmt.Println("Graceful exit")
// 		// 	os.Exit(0)
// 		// }
// 	}()

// 	// go func() {
// 	// 	sigChan := <-sigChan
// 	// 	fmt.Printf("Received signal: %s\n", sigChan)
// 	// 	fmt.Println("Graceful exit")
// 	// 	os.Exit(0)
// 	// }()

// 	// Simulate some work
// 	fmt.Println("Working...")
// 	for {
// 		time.Sleep(time.Second)
// 	}
// }
