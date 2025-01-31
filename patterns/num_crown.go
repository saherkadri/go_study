package patterns

import "fmt"

func Num_Crown() {
	fmt.Println("---- Num_Crown ----")
	n := 5
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		space = space - 2
		fmt.Println()
	}
}

/*
1        1
12      21
123    321
1234  4321
1234554321
*/
