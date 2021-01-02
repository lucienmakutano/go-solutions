package maxsumarr


// MaxSumArr returns an array of two nums when added together make the max sum
func MaxSumArr(nums []int) [2]int {
	if len(nums) == 0 || len(nums) < 2 {
		return [2]int{0,0}
	}

	var res [2]int = [...]int{nums[0], nums[1]}

	for idx, num := range nums {
		if idx < 2 {
			continue
		}

		smIdx := min(res)
		if num > res[smIdx] {
			res[smIdx] = num
		}
	}

	return res
}

func min(nums [2]int) int {
	if nums[0] < nums[1] {
		return 0
	} else {
		return 1
	}

}
