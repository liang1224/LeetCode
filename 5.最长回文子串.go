/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	length := 0

	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)   // 奇
		left2, right2 := expandAroundCenter(s, i, i+1) // 偶

		if right1-left1+1 > length {
			length = right1 - left1 + 1
			start = left1
		}

		if right2-left2+1 > length {
			length = right2 - left2 + 1
			start = left2
		}
	}

	return s[start : start+length]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}

	return left + 1, right - 1
}

// 最长回文子串2
// 动态规划法 时间复杂度:O(n^2) 空间复杂度O(n^2)
func longestPalindrome1(s string) string {
	length := len(s)
	start := 0
	result := 1

	// 构建dp table
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	for l := 0; l < length; l++ {
		for i := 0; i+l < length; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] // 核心: 取之前的计算结果
				}
			}

			// 取最优结果
			if dp[i][j] > 0 && l+1 > result {
				result = l + 1
				start = i
			}
		}
	}

	return s[start : start+result]
}

// @lc code=end