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

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.
//
// https://leetcode.com/problems/unique-number-of-occurrences
//
// Notes: Use a set to count occurences. use another set to check uniqueness of occurence counts.
//
// tags: [hashmap/set, easy]
