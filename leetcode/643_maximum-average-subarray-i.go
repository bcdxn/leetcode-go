package leetcode

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	j := 0
	maxSum := sum

	for j+k < len(nums) {
		sum = sum - nums[j] + nums[j+k]
		if sum > maxSum {
			maxSum = sum
		}
		j++
	}

	return float64(maxSum) / float64(k)
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value and
// return this value. Any answer with a calculation error less than 10-5 will be accepted.
//
// https://leetcode.com/problems/maximum-average-subarray-i
//
// tags: [slidingwindow, easy]
