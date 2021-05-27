/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for next := index + 1; next < len(nums); next++ {
			if target == value+nums[next] {
				return []int{index, next}
			}
		}
	}

	return []int{}
}

// @lc code=end

