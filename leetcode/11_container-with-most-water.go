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
