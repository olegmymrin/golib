package main

import (
	"fmt"
	"time"
)

func publish(events chan string) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Publish before")
		events <- fmt.Sprintf("Hello %d", i)
		fmt.Println("Publish after")
	}
	fmt.Println("Closing channel")
	close(events)
}

func consume(events chan string, end chan bool) {
	for {
		fmt.Println("Consume before")
		e, closed := <-events
		fmt.Printf("Received %s\n", e)
		if closed {
			fmt.Println("Consume on closed")
			break
		}
		fmt.Printf("Consume after %s\n", e)
	}
	end <- true
}

func main() {
	events := make(chan string)
	end := make(chan bool)
	fmt.Println("Begin")
	go publish(events)
	go consume(events, end)
	fmt.Println("End")
	<-end
	time.Sleep(1 * time.Minute)
}
