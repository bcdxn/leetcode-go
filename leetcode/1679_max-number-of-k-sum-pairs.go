package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i := 0
	j := len(nums) - 1
	count := 0

	for i < j {
		if nums[i]+nums[j] == k {
			i++
			j--
			count++
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}

	return count
}
