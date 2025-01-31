package patterns

import "fmt"

func Half_diamond_pattern() {
	fmt.Println("---- Half_diamond_pattern ----")
	n := 5
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* Output

*
**
***
****
*****
*****
****
***
**
*


 */
