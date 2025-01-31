package channels

import (
	"fmt"
	"sync"
	"time"
)

// Producer function
func producer(buffer chan<- int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		item := id*100 + i
		buffer <- item
		fmt.Printf("Producer %d produced %d\n", id, item)
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
}

// Consumer function
func consumer(buffer <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for item := range buffer {
		fmt.Printf("Consumer %d consumed %d\n", id, item)
		time.Sleep(time.Millisecond * 150) // Simulate work
	}
}

const bufferSize = 5

func Producer_Consumer() {
	buffer := make(chan int, bufferSize)
	var producerWg, consumerWg sync.WaitGroup

	// Number of producers and consumers
	numProducers := 2
	numConsumers := 3

	// Start producers
	for i := 1; i <= numProducers; i++ {
		producerWg.Add(1)
		go producer(buffer, &producerWg, i)
	}

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		consumerWg.Add(1)
		go consumer(buffer, &consumerWg, i)
	}

	// Wait for all producers to finish
	producerWg.Wait()
	close(buffer) // Close the buffer channel to signal consumers

	// Wait for all consumers to finish
	consumerWg.Wait()
}
