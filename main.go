package main

import (
	"fmt"
	"go_study/basic_programs"
	"go_study/linked_list"
	"go_study/patterns"
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

	// Fibonacci Series for 10 elements
	fibo_series := basic_programs.Fibonacci(10)
	fmt.Println(fibo_series)

	// Check if number is in fibonacci series
	n := 11
	isfibo := basic_programs.Contains_ele_fibo(fibo_series, n)
	fmt.Printf("fibonacci series contains number %d : %t", n, isfibo)
	fmt.Println()

	// Display n prime numbers
	s := 10
	fmt.Println(basic_programs.Prime_numbers_Series(s))

	// Check if number is prime number
	fmt.Println("---- Num_is_prime ----")
	fmt.Println(basic_programs.Num_is_prime(5))

	// Merge and sort 2 arrays
	array1 := []int{1, 2, 3, 0, 0, 0}
	array2 := []int{2, 5, 6}
	m, n := 3, 3
	basic_programs.Merge(array1, m, array2, n)

	// Remove duplicates
	k := basic_programs.RemoveDuplicates(array1)
	fmt.Println(k)

	// Remove specified element from array
	remove_ele := 3
	res := basic_programs.RemoveElement(array1, remove_ele)
	fmt.Println(res)

	// Pattern Printing
	patterns.Rectangle()
	patterns.Star_pyramid()
	patterns.Right_angled_triangle()
	patterns.Right_angles_same_num_rows()
	patterns.Right_angles_num_pyramid()
	patterns.Num_Crown()
	patterns.Inverter_pyramid()
	patterns.Inverted_right_pyramid()
	patterns.Inverted_right_num_pyramid()
	patterns.Increasing_num_right_triangle()
	patterns.Increasing_decreasing_letter_triangle()
	patterns.Half_diamond_pattern()
	patterns.Diamond()
	patterns.Alternate_binary_num_right_angle()

	// Linked List
	// Creating two linked lists
	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println("Create linked list from array and sum elements and print res i.e 2+5, 4+6,3+4 = 807 starting from last")
	l1 := linked_list.CreateList([]int{2, 4, 3}) // represents the number 342
	l2 := linked_list.CreateList([]int{5, 6, 4}) // represents the number 465
	linked_list.Printll(l1)
	linked_list.Printll(l2)
	// Adding the two numbers
	result := linked_list.AddTwoNumbers(l1, l2)

	// Printing the result linked list
	fmt.Println("... Addition of LinkedList elements ...")
	linked_list.Printll(result) // Output should be 7 -> 0 -> 8 (which is 807)
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
