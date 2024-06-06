package leetcode

import "math/rand/v2"

func maxOperations(nums []int, k int) int {
	quicksort(nums, 0, len(nums)-1)

	i := 0
	j := len(nums) - 1
	count := 0

	for i < j {
		if nums[i]+nums[j] == k {
			i++
			j--
			count++
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}

	return count
}

func quicksort(nums []int, l, r int) {
	if r <= l {
		return
	}
	// pick partition
	p := rand.IntN(r-l+1) + l
	// move partition to the end
	nums[p], nums[r] = nums[r], nums[p]
	// partition
	p = partition(nums, l, r)
	// quicksort left of partition
	quicksort(nums, l, p-1)
	// quicksort right of partition
	quicksort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	// keep track of the first element greater than the partition value
	j := l
	for i := l; i < r; i++ {
		if nums[i] < nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	// swap partition value with the first index where a >= number was found
	nums[j], nums[r] = nums[r], nums[j]
	// return the final resting place of the partition
	return j
}
