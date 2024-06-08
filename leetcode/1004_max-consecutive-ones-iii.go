package leetcode

func longestOnes(nums []int, k int) int {
	c := 0
	mc := 0
	f := 0
	l := 0
	r := 0

	for r < len(nums) {
		if nums[r] == 0 {
			f++
		}

		c++

		if f > k {
			if nums[l] == 0 {
				f--
			}
			l++
			c--
		}

		if f <= k && c > mc {
			mc = c
		}

		r++
	}

	return mc
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the
// array if you can flip at most k 0's.
//
// https://leetcode.com/problems/max-consecutive-ones-iii/description
//
// Notes: kind of a combination of sliding window and two pointers. Move right boundary out to the
// right until end. Move left boundary when we've flipped too many bits. only update max count when
// we haven't flipped too many bits
//
// tags: [slidingwindow, medium, phard]
