package channels

import (
	"fmt"
	"time"
)

func Buffered_chan() {
	fmt.Println("---- Buffered_chan ----")
	mychan := make(chan string, 2)
	mychan <- "saher"
	mychan <- "kadri"
	fmt.Println(<-mychan)
	fmt.Println(<-mychan)
	time.Sleep(time.Second * 2)
}

func Unbuffered_chan() {
	fmt.Println("---- Unbuffered_chan ----")
	// unbuffered channel
	mychan := make(chan int)

	// first read from channel
	go func() {
		fmt.Println(<-mychan)
	}()

	// Attempt to send the value 1 into the channel
	// This operation will block indefinitely because there's no receiver
	// ready to receive the value
	mychan <- 1

	time.Sleep(time.Second * 2)
}
