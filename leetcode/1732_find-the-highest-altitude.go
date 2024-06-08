package leetcode

func largestAltitude(gain []int) int {
	a := 0
	ha := 0

	for _, v := range gain {
		a += v

		if a > ha {
			ha = a
		}
	}

	return ha
}
