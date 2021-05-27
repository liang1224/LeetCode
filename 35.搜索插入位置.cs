/*
 * @lc app=leetcode.cn id=35 lang=csharp
 *
 * [35] 搜索插入位置
 */

// @lc code=start
public class Solution
{
    public int SearchInsert(int[] nums, int target)
    {
        if (nums.Length == 0)
        {
            return 0;
        }

        int s = 0, e = nums.Length;
        var n = e / 2;

        while (e >= s)
        {
            if (target == nums[n])
            {
                return n;
            }

            if (e - s <= 1)
            {
                
            }

            if (e - s)

                if (target > nums[0])
                {
                    s = n;
                    n = (e - n) / 2;
                }
                else
                {
                    e = n;
                    n = (n - s) / 2;
                }
        }
    }
}
// @lc code=end

