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
