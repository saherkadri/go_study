// refernce link
// https://tutorialedge.net/golang/go-linked-lists-tutorial/#:~:text=Let%E2%80%99s%20first%20take%20a%20look%20at%20how%20we,with%20%271%27%20at%20the%20back%20of%20the%20list
// https://www.sibis.dev/posts/linked-lists-in-golang-explained-with-real-world-application/#:~:text=Singly-linked%20list%20maintains%20head%20node%20and%20tail%20node.,tail%20node%20denoted%20the%20end%20of%20the%20list.

package linked_list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func CreateList(myslice []int) *Node {
	fmt.Println("... CreateList ...")
	head := &Node{}
	current := head
	for _, ele := range myslice {
		current.Next = &Node{Value: ele}
		current = current.Next
	}
	fmt.Println("Linked List :", head.Next)
	return head.Next
}

func Printll(head *Node) {
	fmt.Println("... Printing Linked List ...")
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Print(curr.Value, "->")
	}
	fmt.Println(nil)
}

func AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	fmt.Println("... AddTwoNumbers ...")
	// Create a dummy node to simplify handling the result
	dummy := &Node{}
	current := dummy
	fmt.Println("curr:", current)
	carry := 0
	// Traverse both lists until both are exhausted
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		// Add the value of the current node of l1, if it exists
		if l1 != nil {
			fmt.Printf("Adding elements %d", l1.Value)
			//fmt.Println("curr ele:", l1.Value)
			sum = sum + l1.Value
			l1 = l1.Next
		}
		// Add the value of the current node of l2, if it exists
		if l2 != nil {
			fmt.Printf(" + %d", l2.Value)
			//fmt.Println("curr ele:", l2.Value)
			sum = sum + l2.Value
			l2 = l2.Next
		}
		fmt.Println()
		fmt.Println("sum is:", sum)
		// Update the carry for the Next iteration
		carry = sum / 10
		fmt.Println("carry is:", carry)

		// Create a new node for the current digit
		//fmt.Println("sum mod:", sum%10)
		current.Next = &Node{Value: sum % 10}
		fmt.Println("current Next:", current.Next)
		//Move to Next node
		current = current.Next

	}
	// Return the result list (starting from the Next of the dummy node)
	return dummy.Next
}
