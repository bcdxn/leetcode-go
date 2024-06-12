package leetcode

import "strconv"

func equalPairs(grid [][]int) int {
	l := len(grid)
	// store rows in map
	rows := make(map[string]int)
	for r := 0; r < l; r++ {
		key := ""
		for c := 0; c < l; c++ {
			key += strconv.Itoa(grid[r][c]) + ":"
		}
		if _, ok := rows[key]; !ok {
			rows[key] = 0
		}
		rows[key]++
	}
	// count how many columns are in the map
	count := 0
	for c := 0; c < l; c++ {
		key := ""
		for r := 0; r < l; r++ {
			key += strconv.Itoa(grid[r][c]) + ":"
		}
		if v, ok := rows[key]; ok {
			count += v
		}
	}

	return count
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri
// and column cj are equal.
//
// A row and column pair is considered equal if they contain the same elements in the same order
// (i.e., an equal array).
//
// https://leetcode.com/problems/equal-row-and-column-pairs
//
// tags: [hashmap/set, medium]
