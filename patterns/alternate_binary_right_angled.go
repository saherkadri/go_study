package patterns

import "fmt"

func Alternate_binary_num_right_angle() {
	fmt.Println("---- Alternate_binary_num_right_angle ----")
	fmt.Println("... Starting with 0 on top and the alternate start on each row...")
	num := 1
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			num = 1
		} else {
			num = 0
		}
		for j := 0; j < i; j++ {
			fmt.Print(num)
			num = 1 - num
		}
		fmt.Println()
	}
	/*

		0
		10
		010
		1010
		01010

	*/
	fmt.Println()
	fmt.Println()

	// Start with 1 on each row
	fmt.Println("... Starting with 1 on each row ...")
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			num = 1
		}
		for j := 0; j < i; j++ {
			fmt.Print(num)
			num = 1 - num
		}
		fmt.Println()
	}
	/*

		1
		10
		101
		1010
		10101

	*/
	fmt.Println()
	fmt.Println()

	// Start with 0 on each row
	fmt.Println("... Starting with 0 on each row ...")
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			num = 0
		}
		for j := 0; j < i; j++ {
			fmt.Print(num)
			num = 1 - num
		}
		fmt.Println()
	}
	/*

		0
		01
		010
		0101
		01010

	*/
}
