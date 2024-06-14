package leetcode

import "strconv"

func decodeString(s string) string {
	countStack := []int{}
	sequenceStack := [][]byte{}

	sequence := []byte{}
	count := 0
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			count = count*10 + int(s[i]) - '0'
		} else if s[i] == '[' {
			countStack = append(countStack, count)
			sequenceStack = append(sequenceStack, sequence)
			// reset count
			count = 0
			// reset sequence
			sequence = []byte{}
		} else if s[i] == ']' {
			// 'pop' count
			k := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			tmp := sequence
			// 'pop' sequence
			sequence = sequenceStack[len(sequenceStack)-1]
			sequenceStack = sequenceStack[:len(sequenceStack)-1]

			for k > 0 {
				sequence = append(sequence, tmp...)
				k--
			}
		} else {
			sequence = append(sequence, s[i])
		}
	}

	return string(sequence)
}

// # Description:
//
// Given an encoded string, return its decoded string.
//
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is
// being repeated exactly k times. Note that k is guaranteed to be a positive integer.
//
// You may assume that the input string is always valid; there are no extra white spaces, square
// brackets are well-formed, etc. Furthermore, you may assume that the original data does not
// contain any digits and that digits are only for those repeat numbers, k. For example, there will
// not be input like 3a or 2[4].
//
// The test cases are generated so that the length of the output will never exceed 10^5.
//
// https://leetcode.com/problems/decode-string
//
// Notes: Manage 2 stacks, one for the counts, and one for the sub sequences
//
// tags: [stack, medium, phard]
