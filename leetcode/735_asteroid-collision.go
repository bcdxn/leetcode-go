package leetcode

import "math"

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) < 1 {
		return asteroids
	}

	remaining := make([]int, 0, len(asteroids))
	chainReaction := false

	i := 0

	for i < len(asteroids) || (chainReaction && len(remaining) > 1) {
		var left int
		var right int

		if chainReaction && len(remaining) > 1 {
			left = remaining[len(remaining)-2]
			right = remaining[len(remaining)-1]
			remaining = remaining[:len(remaining)-2]
		} else if len(remaining) > 0 {
			left = remaining[len(remaining)-1]
			remaining = remaining[:len(remaining)-1]
			right = asteroids[i]
			i++
		} else if i < len(asteroids)-1 {
			left = asteroids[i]
			i++
			right = asteroids[i]
			i++
		} else if i < len(asteroids) {
			remaining = append(remaining, asteroids[i])
			i++
			continue
		}

		chainReaction = false

		if left > 0 && right < 0 { // collision course
			if left > int(math.Abs(float64(right))) {
				remaining = append(remaining, left)
			} else if left < int(math.Abs(float64(right))) {
				remaining = append(remaining, right)
			} // they're equal and both lose
			chainReaction = true
		} else {
			remaining = append(remaining, left, right)
		}
	}

	return remaining
}

// # Description:
//
// We are given an array asteroids of integers representing asteroids in a row.
//
// For each asteroid, the absolute value represents its size, and the sign represents its direction
// (positive meaning right, negative meaning left). Each asteroid moves at the same speed.
//
// Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one
// will explode. If both are the same size, both will explode. Two asteroids moving in the same
// direction will never meet.
//
// https://leetcode.com/problems/asteroid-collision
//
// Notes: remember chain reaction of collisions
//
// tags: [stack, medium]
