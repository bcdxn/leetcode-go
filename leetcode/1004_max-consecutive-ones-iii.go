package leetcode

func longestOnes(nums []int, k int) int {
	c := 0
	mc := 0
	f := 0
	i := 0
	j := 0

	for i <= j && j < len(nums) {
		for j < len(nums) && f <= k {
			if nums[j] == 0 {
				f++
			}

			c++
			j++

			if f <= k && c > mc {
				mc = c
			}
		}

		for f > k && i <= j {
			if nums[i] == 0 {
				f--
			}
			c--
			i++
		}
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
// right until max bits flipped (or end). Move left boundary in until we have more bits to flip.
// Keeping track of when to update max count was the trick (inside of right boundary movement &&
// ensure we haven't flipped too many bits yet)
//
// tags: [slidingwindow, medium, phard]
