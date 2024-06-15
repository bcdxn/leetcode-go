package leetcode

import "strconv"

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) push(e T) {
	s.s = append(s.s, e)
}

func (s *Stack[T]) pop() T {
	e := s.s[len(s.s)-1]
	s.s = s.s[0 : len(s.s)-1]
	return e
}

func decodeString(s string) string {
	countStack := Stack[int]{}
	sequenceStack := Stack[[]byte]{}

	sequence := []byte{}
	count := 0
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			count = count*10 + int(s[i]) - '0'
		} else if s[i] == '[' {
			countStack.push(count)
			sequenceStack.push(sequence)
			count = 0           // reset count
			sequence = []byte{} // reset sequence
		} else if s[i] == ']' {
			k := countStack.pop()
			tmp := sequence
			sequence = sequenceStack.pop()

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
