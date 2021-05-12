/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	container := make(map[byte]int)
	result := 0
	flag := 0

	for left, right := 0, 0; right < length && left <= right; right++ {
		value, ok := container[s[right]]
		if ok {
			left += value - flag

			for k := flag; k < left; k++ {
				delete(container, s[k])
			}

			flag = value
		}

		container[s[right]] = right + 1

		if len(container) > result {
			result = len(container)
		}
	}

	return result
}

// @lc code=end

