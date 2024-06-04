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
