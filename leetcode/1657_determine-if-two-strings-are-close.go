package leetcode

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	f1 := make([]int, 26)
	f2 := make([]int, 26)

	charSet := make(map[rune]struct{})

	for _, c := range word1 {
		charSet[c] = struct{}{}
		f1[c-'a']++
	}

	for _, c := range word2 {
		// If there are any characters in word2 _not_ in word1 then we can return false
		if _, ok := charSet[c]; !ok {
			return false
		}
		f2[c-'a']++
	}

	sort.Slice(f1, func(i, j int) bool {
		return f1[i] < f1[j]
	})

	sort.Slice(f2, func(i, j int) bool {
		return f2[i] < f2[j]
	})

	for i := 0; i < len(f1); i++ {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}

// all set implementation without ascii arithmetic tric k
// func closeStrings(word1 string, word2 string) bool {
// 	if len(word1) != len(word2) {
// 		return false
// 	}

// 	word1Set := make(map[byte]int)
// 	// word1Counts := make(map[int]struct{})
// 	word2Set := make(map[byte]int)
// 	// word2Counts := make(map[int]struct{})

// 	for i := range len(word1) {
// 		if _, ok := word1Set[word1[i]]; !ok {
// 			word1Set[word1[i]] = 0
// 		}
// 		if _, ok := word2Set[word2[i]]; !ok {
// 			word2Set[word2[i]] = 0
// 		}

// 		word1Set[word1[i]]++
// 		word2Set[word2[i]]++
// 	}

// 	// they must have the same number of letters
// 	if len(word1Set) != len(word2Set) {
// 		return false
// 	}
// 	// they must have the same sets of characters
// 	for k := range word1Set {
// 		if _, ok := word2Set[k]; !ok {
// 			return false
// 		}
// 	}

// 	// Convert the counts to an array and sort them
// 	word1LetterCounts := make([]int, 0, len(word1))
// 	for _, v := range word1Set {
// 		word1LetterCounts = append(word1LetterCounts, v)
// 	}
// 	word2LetterCounts := make([]int, 0, len(word2))
// 	for _, v := range word2Set {
// 		word2LetterCounts = append(word2LetterCounts, v)
// 	}

// 	if len(word1LetterCounts) != len(word2LetterCounts) {
// 		return false
// 	}

// 	sort.Slice(word1LetterCounts, func(i, j int) bool {
// 		return word1LetterCounts[i] < word1LetterCounts[j]
// 	})

// 	sort.Slice(word2LetterCounts, func(i, j int) bool {
// 		return word2LetterCounts[i] < word2LetterCounts[j]
// 	})
// 	// They must have the same counts
// 	for i := range len(word1LetterCounts) {
// 		if word1LetterCounts[i] != word2LetterCounts[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Two strings are considered close if you can attain one from the other using the following operations:
//
// * Operation 1: Swap any two existing characters.
//   For example, abcde -> aecdb
// * Operation 2: Transform every occurrence of one existing character into another existing
//   character, and do the same with the other character.
//   For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
//
// You can use the operations on either string as many times as necessary.

// Given two strings, word1 and word2, return true if word1 and word2 are close, and false
// otherwise.
//
// https://leetcode.com/problems/determine-if-two-strings-are-close
//
// Notes: Can be solved by checking frequency distribution + same characters in each word; order
// doesn't matter. Neat trick by creating array of fixed size (num characters in the alphabet) and
// storing counts at index corresponding to ascii code (ascii code of char - ascii code of 'a' to
// get zero-based index). It can also be solved using multiple sets/maps if you can't think of the
// clever ascii/zero-based index array trick.
//
// tags: [slidingwindow, medium]
