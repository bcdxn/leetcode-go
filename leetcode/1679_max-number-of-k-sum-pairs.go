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

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// You are given an integer array nums and an integer k.
//
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from
// the array.
//
// Return the maximum number of operations you can perform on the array.
//
// https://leetcode.com/problems/max-number-of-k-sum-pairs
//
// tags: [2ptr, medium]
