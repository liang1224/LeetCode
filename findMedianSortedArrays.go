package LeetCode

import "math"

//TODO

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	left := (len1 + len2 + 1) / 2
	right := (len1 + len2 + 2) / 2

	return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / float64(2)
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) float64 {
	if k > len(nums1) {
		return float64(nums2[j+k-1])
	}
	if k > len(nums2) {
		return float64(nums1[i+k-1])
	}

	if k == 1 {
		return math.Min(float64(nums1[i]), float64(nums2[j]))
	}

	mid1, mid2 := 10^6+1, 10^6+1

	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}

	if mid1 < mid2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j-k/2, k-k/2)
	}
}
