package patterns

import "fmt"

func Star_pyramid() {
	fmt.Println("---- Star_pyramid ----")
	n := 5
	for i := 0; i < n; i++ {
		//for adding spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* Output


   *
  ***
 *****
*******

*/
