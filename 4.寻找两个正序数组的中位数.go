/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	total := len1 + len2

	if total == 0 {
		return 0
	}

	if total == 1 {
		if len1 == 1 {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	// 定义指针
	p1, p2 := 0, 0

	// 判断奇偶(能否奇偶使用同一个方法处理？)
	if total%2 == 0 {
		if nums1[p]
	}
}

// @lc code=end