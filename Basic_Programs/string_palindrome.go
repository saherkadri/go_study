package basic_programs

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) {
	fmt.Println("---- IsPalindrome ----")
	// Convert to lowercase to make the check case-insensitive
	s = strings.ToLower(s)
	// Remove spaces (optional, if you want to consider "A man a plan a canal Panama" as a palindrome)
	s = strings.ReplaceAll(s, " ", "")

	// Check if the string is the same forwards and backwards
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			//return false
			fmt.Println("Not Palindrome")
		}
	}
	//return true
	fmt.Println("Palindrome")
}
