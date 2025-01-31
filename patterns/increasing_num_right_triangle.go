package patterns

import "fmt"

func Increasing_num_right_triangle() {
	fmt.Println("---- Increasing_num_right_triangle ----")
	n := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(n)
			n = n + 1
		}
		fmt.Println()
	}
}

/*
1
23
456
78910
*/
