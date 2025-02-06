package basic_programs

import "fmt"

func FindPairs(arr []int, target int) [][2]int {
	fmt.Println(".... FindPairs ....")
	pairs := [][2]int{}
	seen := make(map[int]bool)

	for _, num := range arr {
		complement := target - num
		if seen[complement] {
			pairs = append(pairs, [2]int{complement, num})
		}
		seen[num] = true
	}

	return pairs
}

func Find_pairs_of_ele_equal_to_sum() {
	fmt.Println("---- Find_pairs_of_ele_equal_to_sum ----")
	arr := []int{2, 4, 3, 7, 1, 5, 9, 6}
	target := 12
	result := FindPairs(arr, target)

	fmt.Println("Pairs that sum to", target, "are:", result)
}
