package patterns

import "fmt"

func normal_pyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func inverted_pyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Diamond() {
	fmt.Println("---- Diamond ----")
	n := 5
	normal_pyramid(n)
	inverted_pyramid(n)
}

/* Output

    *
   ***
  *****
 *******
*********
*********
 *******
  *****
   ***
    *

*/
