package leetcode

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	pivot := -1
	// Calculate prefix and suffix sums for each element
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	// sum at zero and suffix at end are 0
	prefix[0] = 0
	prefix[1] = nums[0]
	suffix[len(nums)-1] = 0
	suffix[len(nums)-2] = nums[len(nums)-1]

	for i := 2; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
		suffix[len(nums)-1-i] = suffix[len(nums)-i] + nums[len(nums)-i]
	}

	for i := range nums {
		if prefix[i] == suffix[i] {
			return i
		}
	}

	return pivot
}

// # Description:
//
// Given an array of integers nums, calculate the pivot index of this array.
//
// The pivot index is the index where the sum of all the numbers strictly to the left of the index
// is equal to the sum of all the numbers strictly to the index's right.
//
// If the index is on the left edge of the array, then the left sum is 0 because there are no
// elements to the left. This also applies to the right edge of the array.
//
// Return the leftmost pivot index. If no such index exists, return -1.
//
// https://leetcode.com/problems/find-pivot-index
//
// Notes: Be careful with index arthmetic and write out prefix, suffix, nums arrays in alignment
//
// tags: [prefixsum, easy, pmedium]
