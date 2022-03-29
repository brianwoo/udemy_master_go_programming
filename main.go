/////////////////////////////////
// Mutexes
// Go Playground: https://play.golang.org/p/xVBcUn-CS_4
/////////////////////////////////

//** IMPORTANT **//
// Run this program on your local machine (not in Go Playground)
// Execute: go run main.go

// Use Go Race Detector to check that there is no Data Race
// Execute: go run -race main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func wait(n int, ch chan int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("Waited %d seconds\n", n)
	ch <- n
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func old() {

	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var shared int = 0 // Declaring a shared value

	var m sync.Mutex // 1. Declaring a mutex.

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock() // 2. Lock the access to shared
			shared++
			m.Unlock() // 3. Unlock shared after it's incremented
			wg.Done()
		}()

		// Doing the same for the 2nd goroutine
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			shared--
			wg.Done()
		}()

	}

	wg.Wait()

	// printing the final value of n
	// the final final of n will be always 0
	fmt.Println(shared)
}
