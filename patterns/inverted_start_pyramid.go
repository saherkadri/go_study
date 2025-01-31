package patterns

import "fmt"

func Inverter_pyramid() {
	fmt.Println("---- Inverter_pyramid ----")
	n := 5
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*n-(2*i-1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/* Output

***********
 *********
  *******
   *****
    ***
     *

*/
