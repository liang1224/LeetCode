/*
 * @lc app=leetcode.cn id=21 lang=csharp
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution
{
    public ListNode MergeTwoLists(ListNode l1, ListNode l2)
    {
        if (l1 == null)
        {
            return l2;
        }

        if (l2 == null)
        {
            return l1;
        }

        var root = l1.val > l2.val ? l2 : l1;
        var cursor = root;

        var cursor1 = l1.val > l2.val ? l1 : l1.next;
        var cursor2 = l1.val > l2.val ? l2.next : l2;

        while (cursor1 != null || cursor2 != null)
        {
            if (cursor1 == null)
            {
                cursor.next = cursor2;
                cursor2 = cursor2.next;
                cursor = cursor.next;
                continue;
            }

            if (cursor2 == null)
            {
                cursor.next = cursor1;
                cursor1 = cursor1.next;
                cursor = cursor.next;
                continue;
            }

            if (cursor1.val <= cursor2.val)
            {
                cursor.next = cursor1;
                cursor1 = cursor1.next;
                cursor = cursor.next;
            }
            else
            {
                cursor.next = cursor2;
                cursor2 = cursor2.next;
                cursor = cursor.next;
            }
        }

        return root;
    }
}
// @lc code=end

public class ListNode
{
    public int val;
    public ListNode next;
    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}