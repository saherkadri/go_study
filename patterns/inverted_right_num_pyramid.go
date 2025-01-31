package patterns

import "fmt"

func Inverted_right_num_pyramid() {
	fmt.Println("---- Inverted_right_num_pyramid ----")
	n := 5
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Print(n - j + 1)
		}
		fmt.Println()
	}
}

/* Output

12345
1234
123
12
1

*/
