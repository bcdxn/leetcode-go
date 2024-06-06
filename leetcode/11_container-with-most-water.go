package leetcode

import "math"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	a := area(height, left, right)

	for left < right {
		// move the 'shorter' of the two pointers inward; this would be the only way to grow the area
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		// Check new area resulting in moving the bounding pointer
		newArea := area(height, left, right)
		if newArea > a {
			a = newArea
		}
	}

	return a
}

func area(height []int, left, right int) int {
	return int(math.Min(float64(height[left]), float64(height[right])) * float64(right-left))
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// You are given an integer array height of length n. There are n vertical lines drawn such that the
// two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains
// the most water.
//
// Return the maximum amount of water a container can store. Notice that you may not slant the
// container.
//
// https://leetcode.com/problems/container-with-most-water
//
// # Notes:
//
// Start outside  (maximize width) and work your way in, the only way to possibly grow the area is
// to move the shorter of the two container sides
//
// tags: [2ptr, medium]
