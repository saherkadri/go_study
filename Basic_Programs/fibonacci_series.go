package basic_programs

func Fibonacci(n int) []int {
	fibSeries := make([]int, n)
	if n > 0 {
		fibSeries[0] = 0
	}
	if n > 1 {
		fibSeries[1] = 1
	}

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}
	return fibSeries
}

func Contains_ele_fibo(arr []int, element int) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}
