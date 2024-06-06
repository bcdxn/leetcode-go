package leetcode

import "strconv"

func compress(chars []byte) int {
	l := 0
	i := 0
	j := 0

	for j < len(chars) {
		cc := chars[j]
		count := 0

		for j < len(chars) && chars[j] == cc {
			count++
			j++
		}

		chars[i] = cc
		l++
		i++
		if count > 1 {
			countStr := []byte(strconv.Itoa(count))
			for _, v := range countStr {
				chars[i] = v
				l++
				i++
			}
		}
	}

	return l
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given an array of characters chars, compress it using the following algorithm:
//
// Begin with an empty string s. For each group of consecutive repeating characters in chars:
//
// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately, but instead, be stored in the input
// character array chars. Note that group lengths that are 10 or longer will be split into multiple
// characters in chars.
//
// After you are done modifying the input array, return the new length of the array.
//
// You must write an algorithm that uses only constant extra space.
// https://leetcode.com/problems/string-compression/description
//
// tags: [array/string, easy]
