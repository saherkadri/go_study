package basic_programs

import "fmt"

func Shift_right_by_2(arr []int) {
	fmt.Println("---- Shift_right_by_2 ----")
	length := len(arr)
	for i := 0; i < 2; i++ {
		last := arr[length-1]
		for j := length - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = last
	}
	fmt.Println(arr)
}
