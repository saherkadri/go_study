package patterns

import "fmt"

func Rectangle() {
	fmt.Println("---- Rectangle ----")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*****
*****
*****
*****
*****

=== Code Execution Successful ===
*/
