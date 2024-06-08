package leetcode

// vowels map defined in 345

func maxVowels(s string, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			sum++
		}
	}

	maxVowels := sum

	for j := 0; j+k < len(s); j++ {
		if _, ok := vowels[s[j]]; ok {
			sum--
		}
		if _, ok := vowels[s[j+k]]; ok {
			sum++
		}
		if sum > maxVowels {
			maxVowels = sum
		}
	}

	return maxVowels
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given a string s and an integer k, return the maximum number of vowel letters in any substring of
// s with length k.
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
//
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length
//
// tags: [slidingwindow, medium, peasy]
