package leetcode

func moveZeroes(nums []int) {
	// if the length is 1 or less then any 0s are implicitly in the right place
	if len(nums) < 2 {
		return
	}
	// holds index of the right most non-zero
	var p int

	// Find first 0
	for i := 0; i < len(nums); i++ {
		p = i
		if nums[i] == 0 {
			break
		}
	}

	for j := p + 1; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[p], nums[j] = nums[j], nums[p]
			// increment the partition since we just increased the index of the right-most non-zero
			p++
		}
	}
}
