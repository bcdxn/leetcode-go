package leetcode

func isSubsequence(s string, t string) bool {
	match := ""
	sIndex := 0
	tIndex := 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			match += string(s[sIndex])
			sIndex++
			tIndex++
		} else {
			tIndex++
		}
	}

	return len(match) == len(s)
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting
// some (can be none) of the characters without disturbing the relative positions of the remaining
// characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// https://leetcode.com/problems/is-subsequence
//
// tags: [2ptr, easy]
