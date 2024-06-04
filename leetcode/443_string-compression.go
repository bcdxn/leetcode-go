package leetcode

import "strconv"

func compress(chars []byte) int {
	l := 0
	i := 0
	j := 0

	for j < len(chars) {
		cc := chars[j]
		count := 0

		for j < len(chars) && chars[j] == cc {
			count++
			j++
		}

		chars[i] = cc
		l++
		i++
		if count > 1 {
			countStr := []byte(strconv.Itoa(count))
			for _, v := range countStr {
				chars[i] = v
				l++
				i++
			}
		}
	}

	return l
}
