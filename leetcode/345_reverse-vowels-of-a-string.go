package leetcode

var (
	vowels = map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
)

func reverseVowels(s string) string {
	reversed := []byte(s)
	switchIndex := len(s) - 1

	for i := range len(s) {
		if _, ok := vowels[s[i]]; ok {
			for switchIndex > i {
				if _, ok := vowels[s[switchIndex]]; ok {
					reversed[i] = s[switchIndex]
					reversed[switchIndex] = s[i]
					switchIndex--
					break
				} else {
					switchIndex--
				}
			}
		}
		i++
	}

	return string(reversed)
}
