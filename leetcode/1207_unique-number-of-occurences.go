package leetcode

func uniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)

	for _, v := range arr {
		if _, ok := occurences[v]; !ok {
			occurences[v] = 1
		} else {
			occurences[v]++
		}
	}

	uniqueOccurrences := make(map[int]struct{})

	for _, v := range occurences {
		if _, ok := uniqueOccurrences[v]; ok {
			return false
		} else {
			uniqueOccurrences[v] = struct{}{}
		}
	}

	return true
}
