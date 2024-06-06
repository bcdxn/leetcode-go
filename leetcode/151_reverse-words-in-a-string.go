package leetcode

import "strings"

const (
	space = " "
)

func reverseWords(s string) string {
	words := strings.Split(strings.TrimSpace(s), space)
	reversed := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > 0 {
			reversed = append(reversed, words[i])
		}
	}

	return strings.Join(reversed, space)
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will be separated by at
// least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between two words. The
// returned string should only have a single space separating the words. Do not include any extra
// spaces.
//
// https://leetcode.com/problems/reverse-words-in-a-string
//
// # Notes:
//
// Could use some kind of data structure like a stack to reverse (or simply split and iterate in
// reverse)
//
// tags: [string/array, medium]
