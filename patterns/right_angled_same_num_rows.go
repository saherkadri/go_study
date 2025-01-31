package patterns

import "fmt"

func Right_angles_same_num_rows() {
	fmt.Println("---- Right_angles_same_num_rows ----")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

/* Output

1
22
333
4444
55555

*/
