// create a go routine to create stack of 3 ele and pop last element
package channels

import (
	"fmt"
	"sync"
	"time"
)

type Stack struct {
	elements []int
	lock     sync.Mutex
}

func (s *Stack) Push(ele int) {
	fmt.Println(".... Push Element....", ele)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, ele)
}

func (s *Stack) Pop() (int, []int) {
	fmt.Println(".... Pop Element....")
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.elements) == 0 {
		fmt.Println("Empty stack")
		return -1, s.elements
	}
	// len is 4 so last index would be 3 starting from 0
	element := s.elements[len(s.elements)-1]
	// create slice to display elements till 0-3 index
	s.elements = s.elements[:len(s.elements)-1]
	return element, s.elements
}

func Mystack() {
	fmt.Println("---- Execute go routine to pop last element from Mystack ----")
	stack := &Stack{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
	}()
	// Wait for the push goroutine to finish
	wg.Wait()
	// Pop the last element from the stack in the main goroutine
	poppedElement, myslice := stack.Pop()
	fmt.Println("Popped element:", poppedElement)
	fmt.Println("Array after pop:", myslice)
}

func MyApproach_Mystack(myarr []int) {
	var new_arr []int
	for i := 0; i < len(myarr)-1; i++ {
		new_arr = append(new_arr, myarr[i])
	}
	time.Sleep(time.Second * 2)
	fmt.Println(new_arr)
}
