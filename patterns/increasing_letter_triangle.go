package patterns

import "fmt"

func Increasing_decreasing_letter_triangle() {
	fmt.Println("---- Increasing_decreasing_letter_triangle ----")
	n := 5
	/*
		A
		AB
		ABC
		ABCD
		ABCDE
		ABCDEF
	*/
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A' + j))
		}
		fmt.Println()
	}

	/////

	/*
		ABCDEF
		ABCDE
		ABCD
		ABC
		AB
		A
	*/
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			fmt.Print(string('A' + j))
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	////
	/*
		A
		BB
		CCC
		DDDD
		EEEEE
		FFFFFF
		FFFFFF
		EEEEE
		DDDD
		CCC
		BB
		A
	*/

	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A' + i))
		}
		fmt.Println()
	}
	for i := n; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A' + i))
		}
		fmt.Println()
	}

}
