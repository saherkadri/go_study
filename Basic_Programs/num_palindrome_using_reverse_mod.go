package basic_programs

import "fmt"

func IsPalindromeNumber_rev(n int) {
	reverse := 0
	fmt.Println("---- IsPalindromeNumber_rev ----")
	temp := n
	// For Loop used
	for {
		remainder := n % 10 // gives the last digit of a number
		reverse = reverse*10 + remainder
		n = n / 10 //n /= 10
		// fmt.Println("n", n) // after removing last digit displays remaining digits
		if n == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		fmt.Printf("Number %d is a Palindrome", temp)
	} else {
		fmt.Printf("Number %d is not a Palindrome", temp)
	}
	fmt.Println()
}
