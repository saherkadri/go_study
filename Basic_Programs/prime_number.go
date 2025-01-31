package basic_programs

import "fmt"

func Num_is_prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Prime_numbers_Series(n int) []int {
	fmt.Println("---- Prime_numbers_Series ----")
	var prime_numbers []int
	for i := 0; i <= n; i++ {
		if Num_is_prime(i) {
			prime_numbers = append(prime_numbers, i)
		}
	}
	return prime_numbers
}
