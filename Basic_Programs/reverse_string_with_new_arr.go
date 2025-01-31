package basic_programs

import "fmt"

func Rev_arr_with_new_arr(myarr []int) {
	fmt.Println("---- Rev_arr_with_new_arr ----")
	var rev_arr []int
	for i := len(myarr) - 1; i > 0; i-- {
		rev_arr = append(rev_arr, myarr[i])
	}
	fmt.Println(rev_arr)
}
