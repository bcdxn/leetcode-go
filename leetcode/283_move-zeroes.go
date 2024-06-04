package leetcode

func moveZeroes(nums []int) {
	// if the length is 1 or less then any 0s are implicitly in the right place
	if len(nums) < 2 {
		return
	}
	// holds index of the right most non-zero
	partition := -1

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			partition++
			nums[partition], nums[j] = nums[j], nums[partition]
		}
	}

}
