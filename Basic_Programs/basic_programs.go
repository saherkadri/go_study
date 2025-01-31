package basic_programs

import (
	"fmt"
	"strconv"
	"strings"
)

func Rev_arr_with_new_arr(myarr []int) {
	fmt.Println("---- Rev_arr_with_new_arr ----")
	var rev_arr []int
	for i := len(myarr) - 1; i > 0; i-- {
		rev_arr = append(rev_arr, myarr[i])
	}
	fmt.Println(rev_arr)
}

func Rev_arr_without_new_arr(myarr []int) {
	fmt.Println("---- Rev_arr_without_new_arr ----")
	left := 0
	right := len(myarr) - 1
	for left < right {
		myarr[left], myarr[right] = myarr[right], myarr[left]
		left++
		right--
	}
	fmt.Println(myarr)
}

func ResverseStrUsingRune(mystr string) {
	fmt.Println("---- ResverseStrUsingRune ----")
	runes := []rune(mystr)
	for i := len(runes) - 1; i >= 0; i-- {
		//fmt.Printf("%c ", runes[i])
		fmt.Print(string(runes[i]))
	}
	fmt.Println()
}

func Swap2num_operator(numb1 int, numb2 int) {
	fmt.Println("---- Swap2num_operator ----")
	numb1 = numb1 + numb2
	numb2 = numb1 - numb2
	numb1 = numb1 - numb2
	fmt.Println("Swap 2 numbers using operator")
	fmt.Printf("numb1 %d and numb2 %d", numb1, numb2)
	fmt.Println()
}

func Swap2num_xor(numb1 int, numb2 int) {
	fmt.Println("---- Swap2num_xor ----")
	numb1 = numb1 ^ numb2
	numb2 = numb1 ^ numb2
	numb1 = numb1 ^ numb2
	fmt.Println("Swap 2 numbers using xor")
	fmt.Printf("numb1 %d and numb2 %d", numb1, numb2)
	fmt.Println()
}

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

func Shift_right_by_2(arr []int) {
	fmt.Println("---- Shift_right_by_2 ----")
	length := len(arr)
	for i := 0; i < 2; i++ {
		last := arr[length-1]
		for j := length - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = last
	}
	fmt.Println(arr)
}

func Move_zeros_to_right_end(arr []int) {
	fmt.Println("---- Move_zeros_to_right_end ----")
	indx := 0
	for i, _ := range arr {
		if arr[i] != 0 {
			arr[indx], arr[i] = arr[i], arr[indx]
			indx++
		}
	}
	fmt.Println(arr)
}

func Max_ele_in_arr(myarr []int) {
	fmt.Println("---- Max_ele_in_arr ----")
	max := 0
	for i := 0; i < len(myarr); i++ {
		if myarr[i] > max {
			max = myarr[i]
		}
	}
	fmt.Println(max)
}
