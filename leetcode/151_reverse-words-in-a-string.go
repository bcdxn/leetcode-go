package leetcode

import "strings"

const (
	space = " "
)

func reverseWords(s string) string {
	words := strings.Split(strings.TrimSpace(s), space)
	var reversed []string

	for i := range len(words) {
		if len(words[i]) > 0 {
			reversed = append([]string{words[i]}, reversed...)
		}
	}

	return strings.Join(reversed, space)
}
