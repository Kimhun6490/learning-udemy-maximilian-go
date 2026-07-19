package main

import (
	"fmt"
	"time"
)

func greet(name string, done chan bool) {
	fmt.Printf("Hello, %s!\n", name)
	done <- true // Signal that the goroutine is done
}

func slowGreet(name string, done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello, %s!\n", name)
	done <- true // Signal that the goroutine is done
	close(done)  // Close the channel to signal that no more values will be sent
}

func main() {
	// dones := make([]chan bool, 4)
	// for i := range dones {
	// 	dones[i] = make(chan bool)
	// }

	// go greet("Step 1", dones[0])
	// go greet("Step 2", dones[1])
	// go slowGreet("Step 3", dones[2])
	// go greet("Step 4", dones[3])

	// for i := range dones {
	// 	<-dones[i] // Wait for each goroutine to finish
	// }

	done := make(chan bool)
	go greet("Step 1", done)
	go greet("Step 2", done)
	go slowGreet("Step 3", done)
	go greet("Step 4", done)

	for range done {

	}
}
