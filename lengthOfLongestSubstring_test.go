package LeetCode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcdbcgheac"

	result := lengthOfLongestSubstring(s)

	if result != 7 {
		t.Error(result)
	}
}
