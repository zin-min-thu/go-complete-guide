package main

import (
	"fmt"
	"time"
)

func greet(pharse string) {
	fmt.Println("Hello!, ", pharse)
}

func slowGreet(pharse string) {
	time.Sleep(3 * time.Second) // simulate a slow, long taking tasks
	fmt.Println("hello! ", pharse)
}

func main() {
	go greet("Nice to meet you")
	go greet("how are you")
	go slowGreet("How....are...you...?")
	go greet("You are linking the course!")
}
