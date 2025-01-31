package basic_programs

import (
	"fmt"
)

func Max_ele_in_arr(myarr []int) {
	fmt.Println("---- Max_ele_in_arr ----")
	max := 0
	for i := 0; i < len(myarr); i++ {
		if myarr[i] > max {
			max = myarr[i]
		}
	}
	fmt.Println(max)
}
