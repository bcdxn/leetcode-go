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
