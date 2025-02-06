package channels

import (
	"fmt"
	//"sort"
	"sync"
	"time"
)

func even(mychan chan int, myevenwg *sync.WaitGroup) {
	defer myevenwg.Done()
	//mychan <- 2
	for num := range mychan {
		fmt.Println("Even:", num)
		time.Sleep(time.Second * 1)
	}
}

func odd(mychan chan int, myoddwg *sync.WaitGroup) {
	defer myoddwg.Done()
	//mychan <- 3
	for num := range mychan {
		fmt.Println("Odd:", num)
		time.Sleep(time.Second * 1)
	}
}

func Even_Odd_Routines() {
	myEvenchan := make(chan int)
	myOddchan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go even(myEvenchan, &wg)
	go odd(myOddchan, &wg)

	myEvenchan <- 0
	myOddchan <- 0
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			myEvenchan <- i
		} else {
			myOddchan <- i
		}
	}
	close(myEvenchan)
	close(myOddchan)
	wg.Wait()
}
