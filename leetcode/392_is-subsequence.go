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
