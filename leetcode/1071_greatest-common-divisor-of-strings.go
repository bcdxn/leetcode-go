package leetcode

import "math"

func gcdOfStrings(str1 string, str2 string) string {
	divisor := ""
	l1 := len(str1)
	l2 := len(str2)
	minLen := int(math.Min(float64(l1), float64(l2)))
	maxLen := int(math.Max(float64(l1), float64(l2)))
	// must be a common prefix
	for i := 0; i < minLen; i++ {
		if str1[i] == str2[i] {
			divisor += string(str1[i])
		} else {
			continue
		}
	}
	// must be evenly divisible by the devisor
	for len(divisor) > 0 && (minLen%len(divisor) != 0 || maxLen%len(divisor) != 0) {
		divisor = divisor[0 : len(divisor)-1]
	}
	if len(divisor) < 1 {
		return divisor
	}
	// both strings must be made entirely of the divisor
	for i := 0; i < len(str1); i++ {
		if str1[i] != divisor[i%len(divisor)] {
			return ""
		}
	}
	// both strings must be made entirely of the divisor
	for i := 0; i < len(str2); i++ {
		if str2[i] != divisor[i%len(divisor)] {
			return ""
		}
	}

	return divisor
}
