package basic_programs

import "fmt"

func Move_zeros_to_right_end(arr []int) {
	fmt.Println("---- Move_zeros_to_right_end ----")
	indx := 0
	for i, _ := range arr {
		if arr[i] != 0 {
			arr[indx], arr[i] = arr[i], arr[indx]
			indx++
		}
	}
	fmt.Println(arr)
}
