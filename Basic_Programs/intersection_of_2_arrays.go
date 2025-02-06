package basic_programs

import (
	"fmt"
	"sort"
)

func Contains(arr []int, t int) bool {
	fmt.Println("---- Contains ----")
	for _, a := range arr {
		if a == t {
			return true
		}
	}
	return false
}

func Intersection(nums1 []int, nums2 []int) []int {
	fmt.Println("---- Intersection ----")
	flags := map[int]bool{}
	for _, n := range nums1 {
		flags[n] = true
	}

	var intersection []int
	for _, m := range nums2 {
		//if flags[m] && !Contains(intersection, m) {
		if flags[m] {
			intersection = append(intersection, m)
		}
	}
	return intersection
}

func Intersect_elements_two_array() {
	fmt.Println("... Intersect_elements_two_array ...")
	arr1 := []int{2, 4, 6, 8, 10, 12}
	arr2 := []int{3, 6, 9, 12}
	fmt.Println(Intersection(arr1, arr2))
}

func Duplicate(nums []int) {
	dup := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dup = append(dup, nums[i])
		}
	}
	fmt.Println("duplicate :", dup)
}

func Find_duplicate_two_arr() {
	fmt.Println("---- Find_duplicate_two_arr ----")
	arr1 := []int{2, 4, 6, 8, 10, 12}
	arr2 := []int{3, 6, 6, 9, 12}
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	Duplicate(arr1)
}
