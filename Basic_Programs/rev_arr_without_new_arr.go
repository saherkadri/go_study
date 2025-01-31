package basic_programs

import "fmt"

func Rev_arr_without_new_arr(myarr []int) {
	fmt.Println("---- Rev_arr_without_new_arr ----")
	left := 0
	right := len(myarr) - 1
	for left < right {
		myarr[left], myarr[right] = myarr[right], myarr[left]
		left++
		right--
	}
	fmt.Println(myarr)
}
