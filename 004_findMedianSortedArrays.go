package LeetCode

import "math"

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	total := len1 + len2
	isI := true
	i, j := 0, 0

	for i+j <= total/2+1 && i < len1 && j < len2 {
		if len1-i == 0 {
			j++
			isI = false
		}
		if len2-j == 0 {
			i++
			isI = true
		}

		if nums1[i] < nums2[j] {
			i++
			isI = true
		} else {
			j++
			isI = false
		}
	}

	if total%2 == 1 {
		if isI {
			return float64(nums1[i])
		} else {
			return float64(nums2[j])
		}
	} else {
		if isI {
			return (float64(nums1[i]) + math.Min(float64(nums1[i+1]), float64(nums2[j]))) / float64(2)
		} else {
			return (float64(nums2[j]) + math.Min(float64(nums2[j+1]), float64(nums1[i]))) / float64(2)
		}
	}
}
