package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}

	c := math.MaxInt // smallest number we've encountered that has a smaller number
	s := nums[0]     // smallest number we've encountered

	for i := 1; i < l; i++ {
		// check if triplet is satisfied
		if nums[i] > c {
			return true
		}
		// triplet not satisfied; we may update s and/or c

		// update smallest number with a number smaller before it
		if s < nums[i] {
			c = int(math.Min(float64(c), float64(nums[i])))
		}

		// update smallest number encountered
		if nums[i] < s {
			s = nums[i]
		}
	}

	return false
}
