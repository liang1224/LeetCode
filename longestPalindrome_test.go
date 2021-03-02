package LeetCode

import "testing"

func TestLongestPalindrome1(t *testing.T) {
	s := "abcdeftgdgtcace"
	ans := longestPalindrome1(s)

	if "tgdgt" != ans {
		t.Error(ans)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := "abcdeftgdgtcace"
	ans := longestPalindrome2(s)

	if "tgdgt" != ans {
		t.Error(ans)
	}
}
