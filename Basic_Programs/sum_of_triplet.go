package basic_programs

func FindTriplets(arr []int, target int) [][]int {
	var triplets [][]int
	//sort.Ints(arr) // Sort the array

	for i := 0; i < len(arr)-2; i++ {
		// // Avoid duplicate triplets
		// if i > 0 && arr[i] == arr[i-1] {
		//     continue
		// }
		left, right := i+1, len(arr)-1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				triplets = append(triplets, []int{arr[i], arr[left], arr[right]})

				// Avoid duplicates for the second element
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				// Avoid duplicates for the third element
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return triplets
}
