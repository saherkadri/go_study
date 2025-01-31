package basic_programs

import "fmt"

func Swap2num_xor(numb1 int, numb2 int) {
	fmt.Println("---- Swap2num_xor ----")
	numb1 = numb1 ^ numb2
	numb2 = numb1 ^ numb2
	numb1 = numb1 ^ numb2
	fmt.Println("Swap 2 numbers using xor")
	fmt.Printf("numb1 %d and numb2 %d", numb1, numb2)
	fmt.Println()
}
