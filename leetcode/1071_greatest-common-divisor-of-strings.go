package leetcode

// ABCABC ABC
// ABCABCABC+ABC == ABC+ABCABCABC
// This fact must be true for there to be a repeated pattern
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	// Since there is guaranteed to be a repeated pattern we can simply return the greatest
	// common devisor
	minStr := str1
	if len(str2) < len(minStr) {
		minStr = str2
	}

	gcd := len(minStr)
	for gcd > 0 {
		if len(str1)%gcd == 0 && len(str2)%gcd == 0 {
			break
		}
		gcd--
	}

	return minStr[:gcd]
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e.,
// t is concatenated with itself one or more times).
//
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and
// str2.
//
// https://leetcode.com/problems/greatest-common-divisor-of-strings
//
// Notes: This one has a trick - if str1 + str2 == str2 + str1 we can simply find the gcd of the
// length of the strings; if str1 + str2 != str2 + str1 then there is no repeating pattern and
// therefore no gcd pattern (we return an empty string)
//
// tags: [array/string, easy, phard]
