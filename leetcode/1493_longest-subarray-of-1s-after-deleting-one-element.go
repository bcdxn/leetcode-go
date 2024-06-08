package leetcode

const max_removals = 1

func longestSubarray(nums []int) int {
	c := 0
	mc := 0
	l := 0
	r := 0
	removed := 0

	for r < len(nums) {
		if nums[r] == 0 {
			removed++
		} else {
			c++
		}

		if removed > max_removals {
			if nums[l] == 0 {
				removed--
			} else if c > 0 {
				c--
			}
			l++
		}

		if removed == max_removals && c > mc {
			mc = c
		}

		r++
	}

	// We were told we must remove one element; if we have not done so in order to create the longest
	// subarray (by removing a zero), then we can simply remove one from the end and decrement our
	// max count.
	if removed == 0 && c > 0 {
		mc = c - 1
	}

	return mc
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given a binary array nums, you should delete one element from it.
//
// Return the size of the longest non-empty subarray containing only 1's in the resulting array.
// Return 0 if there is no such subarray.
//
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element
//
// Notes: The sliding window _must_ include a zero that is removed (or we need to remove an
// arbitrary 1 from the subarray if it was made without removing a zero)
//
// tags: [slidingwindow, medium]
