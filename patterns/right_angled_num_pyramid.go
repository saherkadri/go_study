package patterns

import "fmt"

func Right_angles_num_pyramid() {
	fmt.Println("---- Right_angles_num_pyramid ----")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

/* Output:
1
12
123
1234
12345
*/
