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
	i := 0

	for i < switchIndex {
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

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases,
// more than once.
//
// https://leetcode.com/problems/reverse-vowels-of-a-string
//
// tags: [array/string, easy]
