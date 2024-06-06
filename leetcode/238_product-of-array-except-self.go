package leetcode

func productExceptSelf(nums []int) []int {
	l := len(nums)
	result := make([]int, l)
	prefixProducts := make([]int, l)
	suffixProducts := make([]int, l)

	prefixProducts[0] = 1
	suffixProducts[l-1] = 1

	for i := 1; i < l; i++ {
		prefixProducts[i] = nums[i-1] * prefixProducts[i-1]
		suffixProducts[l-i-1] = nums[l-i] * suffixProducts[l-i]
	}

	for i := range len(result) {
		result[i] = prefixProducts[i] * suffixProducts[i]
	}

	return result
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given an integer array nums, return an array answer such that answer[i] is equal to the product
// of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// # Notes
//
// Could use some kind of data structure like a stack to reverse (or simply split and iterate in
// reverse)
//
// https://leetcode.com/problems/product-of-array-except-self
//
// tags: [array/string, medium]
