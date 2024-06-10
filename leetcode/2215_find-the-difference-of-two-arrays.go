package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Distinct := make([]int, 0, len(nums2))
	nums2Distinct := make([]int, 0, len(nums2))

	nums1Map := make(map[int]struct{})
	nums1DistinctMap := make(map[int]struct{})
	nums2Map := make(map[int]struct{})
	nums2DistinctMap := make(map[int]struct{})

	for _, v := range nums1 {
		nums1Map[v] = struct{}{}
	}

	for _, v := range nums2 {
		nums2Map[v] = struct{}{}
	}

	for _, v := range nums1 {
		_, inNums2 := nums2Map[v]
		_, inDistinct := nums1DistinctMap[v]
		if !inNums2 && !inDistinct {
			nums1DistinctMap[v] = struct{}{}
			nums1Distinct = append(nums1Distinct, v)
		}
	}

	for _, v := range nums2 {
		_, inNums1 := nums1Map[v]
		_, inDistinct := nums2DistinctMap[v]
		if !inNums1 && !inDistinct {
			nums2DistinctMap[v] = struct{}{}
			nums2Distinct = append(nums2Distinct, v)
		}
	}

	return [][]int{nums1Distinct, nums2Distinct}
}
