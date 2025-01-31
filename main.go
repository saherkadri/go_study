package main

import (
	"fmt"
	"go_study/basic_programs"
)

func main() {
	fmt.Println("Hello World !")

	// Reverse array putting in new arr
	myarr := []int{2, 4, 6, 8, 10, 12, 14, 16}
	basic_programs.Rev_arr_with_new_arr(myarr)

	myarr1 := []int{2, 50, 6, 23, 10, 12, 14, 16}
	basic_programs.Rev_arr_without_new_arr(myarr1)

	// Reverse string with rune
	basic_programs.ResverseStrUsingRune("saherkadri")

	// Swap 2 numbers using operator
	a := 5
	b := 10
	basic_programs.Swap2num_operator(a, b)
	basic_programs.Swap2num_xor(a, b)

	// Check number is Palindrome using strconv
	basic_programs.IsPalindromeNumber_strconv(1234321)

	// Check number is Palindrome usinf rev
	basic_programs.IsPalindromeNumber_rev(1234321)

	// Check string palindrome
	basic_programs.IsPalindrome("Madam")

	// Rotate array elements to right by 2
	arr := []int{1, 2, 3, 4, 5, 6}
	basic_programs.Shift_right_by_2(arr)

	// Move zeros to right end
	arr1 := []int{1, 0, 2, 3, 0, 0, 4, 0, 5}
	basic_programs.Move_zeros_to_right_end(arr1)

	// Find max ele in array

	myarr2 := []int{2, -23, 5, 1, 13}
	basic_programs.Max_ele_in_arr(myarr2)

}

// SA Interview
/*
func getDigit() int {
	return 5
}

func main() {
	//mychan := make(chan int)
	go getDigit()
	fmt.Println(getDigit())
}
*/
