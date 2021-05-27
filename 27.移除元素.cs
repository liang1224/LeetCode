/*
 * @lc app=leetcode.cn id=27 lang=csharp
 *
 * [27] 移除元素
 */

// @lc code=start
public class Solution
{
    public int RemoveElement(int[] nums, int val)
    {
        if (nums.Length == 0)
        {
            return 0;
        }

        int right = nums.Length, left = 0;

        while (left < right)
        {
            if (nums[left] == val)
            {
                nums[left] = nums[right - 1];
                right--;
            }
            else
            {
                left++;
            }
        }

        return left;
    }
}
// @lc code=end

