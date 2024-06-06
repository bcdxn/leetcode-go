package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	planted := 0

	for i := range flowerbed {
		if flowerbed[i] == 0 {
			if (i-1 < 0 || flowerbed[i-1] == 0) && (i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				planted++
			}
		}
	}

	return planted >= n
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// You have a long flowerbed in which some of the plots are planted, and some are not. However,
// flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not
// empty, and an integer n, return true if n new flowers can be planted in the flowerbed without
// violating the no-adjacent-flowers rule and false otherwise.
//
// https://leetcode.com/problems/can-place-flowers/description
//
// tags: [array/string, easy]
