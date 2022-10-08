package main

import (
	"fmt"
)
func main() {

	ans := binarysearch([]int{-1,0,3,5,9,12}, 2)
	fmt.Printf("%v", ans)

}

// find pivot
// if nums[pivot] < target; 

func binarysearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var pivot int
	result := -1
	for start <= end {
		pivot = (end + start) / 2
		if nums[pivot] < target {
			start = pivot + 1
		} else if nums[pivot] > target {
			end = pivot - 1
		} else {
			result = pivot
			break
		}
	}
	return result
}