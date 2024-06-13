package leetcode

func removeStars(s string) string {
	l := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			// don't add the star and remove the previously added character
			l = l[:len(l)-1]
		} else {
			l = append(l, s[i])
		}
	}

	return string(l)
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// You are given a string s, which contains stars *.
//
// In one operation, you can:
//
// - Choose a star in s.
// - Remove the closest non-star character to its left, as well as remove the star itself.
//
// Return the string after all stars have been removed.
// Note:
// The input will be generated such that the operation is always possible.
// It can be shown that the resulting string will always be unique.
//
// https://leetcode.com/problems/removing-stars-from-a-string
//
// tags: [stack, medium,  peasy]
