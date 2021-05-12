/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := ListNode{
		Val:  0,
		Next: nil,
	}
	temp := &root

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0

		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		sum += carry
		carry = 0

		carry = sum / 10
		sum = sum % 10

		temp.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		temp = temp.Next

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	return root.Next
}
// @lc code=end

