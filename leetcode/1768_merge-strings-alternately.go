package leetcode

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	w1Index := 0
	w2Index := 0
	l1 := len(word1)
	l2 := len(word2)
	out := make([]string, 0, l1+l2)

	for w1Index < l1 && w2Index < l2 {
		out = append(out, string(word1[w1Index]))
		w1Index++
		out = append(out, string(word2[w2Index]))
		w2Index++
	}

	for w1Index < l1 {
		out = append(out, string(word1[w1Index]))
		w1Index++
	}

	for w2Index < l2 {
		out = append(out, string(word2[w2Index]))
		w2Index++
	}

	return strings.Join(out, "")
}
