/*
 * @lc app=leetcode.cn id=14 lang=csharp
 *
 * [14] 最长公共前缀
 */

// @lc code=start
public class Solution
{
    public string LongestCommonPrefix(string[] strs)
    {
        string prefix = "";

        for (int j = 0; j < strs[0].Length; j++)
        {
            if (string.IsNullOrWhiteSpace(strs[0]))
            {
                return "";
            }
            prefix += strs[0].Substring(j, 1);

            for (int i = 1; i < strs.Length; i++)
            {
                if (string.IsNullOrWhiteSpace(strs[i]))
                {
                    return "";
                }

                if (!strs[i].StartsWith(prefix))
                {
                    return prefix.Substring(0, prefix.Length - 1);
                }
            }
        }

        return prefix;
    }
}
// @lc code=end

