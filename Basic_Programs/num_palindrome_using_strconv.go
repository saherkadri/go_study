package basic_programs

import (
	"fmt"
	"strconv"
)

func IsPalindromeNumber_strconv(n int) {
	fmt.Println("---- IsPalindromeNumber_strconv ----")
	// Convert the number to a string
	s := strconv.Itoa(n)

	// Check if the string representation of the number is a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			//return false
			fmt.Println("Not Palindrome")
		}
	}
	//return true
	fmt.Println("Palindrome")
}
