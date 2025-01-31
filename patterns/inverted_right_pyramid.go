package patterns

import "fmt"

func Inverted_right_pyramid() {
	fmt.Println("---- Inverted_right_pyramid ----")
	for i := 1; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* Output

*****
****
***
**
*

 */
