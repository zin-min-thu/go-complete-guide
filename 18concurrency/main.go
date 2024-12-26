package main

import (
	"fmt"
	"time"
)

func greet(pharse string, doneChan chan bool) {
	fmt.Println("Hello!, ", pharse)
	doneChan <- true
}

func slowGreet(pharse string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long taking tasks
	fmt.Println("hello! ", pharse)
	doneChan <- true
	close(doneChan) // close can use if you know which channel is taking long tasks
}

func main() {
	// dones := make([]chan bool, 4) // multi custom channel
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you", dones[0])
	// dones[1] = make(chan bool)
	// go greet("how are you", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How....are...you...?", dones[2])
	// dones[3] = make(chan bool)
	// go greet("You are linking the course!", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	done := make(chan bool) // single channel

	go greet("Nice to meet you", done)
	go greet("how are you", done)
	go slowGreet("How....are...you...?", done)
	go greet("You are linking the course!", done)

	for range done {

	}

}
