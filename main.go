package LeetCode

// 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	container := make(map[byte]int)
	array := []byte(s)
	result := 0
	flag := 0

	for left, right := 0, 0; right < length && left <= right; right++ {
		value, ok := container[array[right]]
		if ok {
			left += value - flag

			for k := flag; k < left; k++ {
				delete(container, array[k])
			}

			flag = value
		}

		container[array[right]] = right + 1

		if len(container) > result {
			result = len(container)
		}
	}

	return result
}
