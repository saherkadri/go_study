package patterns

import "fmt"

func Right_angled_triangle() {
	fmt.Println("---- Right_angled_triangle ----")
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Output:

*
**
***
****
*****

*/
