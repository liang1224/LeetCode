package LeetCode

// 两数之和
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