package basic_programs

import "fmt"

func ResverseStrUsingRune(mystr string) {
	fmt.Println("---- ResverseStrUsingRune ----")
	runes := []rune(mystr)
	for i := len(runes) - 1; i >= 0; i-- {
		//fmt.Printf("%c ", runes[i])
		fmt.Print(string(runes[i]))
	}
	fmt.Println()
}
