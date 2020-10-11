package main

import (
	"fmt"
	"math"
)

func main() {
	ok1 := TwoNumbersSumToK([]int{1, 3, 8, 2}, 10)
	ok2 := TwoNumbersSumToK([]int{3, 9, 13, 7}, 8)
	ok3 := TwoNumbersSumToK([]int{4, 2, 6, 5, 2}, 4)

	fmt.Printf("1: %v \n2: %v \n3: %v", ok1, ok2, ok3)
}

// TwoNumbersSumToK return true if 2 nums from the nums slice sums up to k
func TwoNumbersSumToK(nums []int, k int) bool {

	if len(nums) == 0 {
		return false
	}

	for i := range nums {
		sub := int(math.Abs(float64(i - k)))

		for j := range nums {
			if j == sub {
				return true
			}
		}
	}

	return false
}
