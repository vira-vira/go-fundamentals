package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency:
// Multiple processes executing independently
// Lots of tasks "on the go", but only one active

// Go implements goroutine for concurrency, not uses the OS for threads
// - Scheduled by Go runtime
// - Grow and shrink
// - Go manages goroutines
// - Fewer context switches
// - Use less threads
// - Fast startup times
// - Communicate via channels

// Go's Concurrency model:
// - Actor
// - Communicating Sequential Processes (CSP)

// Channel are like pipes

// Without concurrency:
//func main() {
//
//	func() {
//		time.Sleep(5 * time.Second)
//		fmt.Println("Hola")
//	}()
//
//	func() {
//		fmt.Println("Manola")
//	}()
//}

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hola")
	}()

	func() {
		defer waitGrp.Done()
		fmt.Println("Manola")
	}()
	waitGrp.Wait()
}
