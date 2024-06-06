package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// find max
	max := 0

	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	// Create boolean array based on the extra candies
	output := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= max {
			output[i] = true
		}
	}

	return output
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// There are n kids with candies. You are given an integer array candies, where each candies[i]
// represents the number of candies the ith kid has, and an integer extraCandies, denoting the
// number of extra candies that you have.
//
// Return a boolean array result of length n, where result[i] is true if, after giving the ith kid
// all the extraCandies, they will have the greatest number of candies among all the kids, or false
// otherwise.
//
// Note that multiple kids can have the greatest number of candies.
//
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
//
// tags: [array/string, easy]
