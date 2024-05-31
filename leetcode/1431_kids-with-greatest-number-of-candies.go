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
