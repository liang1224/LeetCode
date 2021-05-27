/*
 * @lc app=leetcode.cn id=13 lang=csharp
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
public class Solution
{
    static Dictionary<char, int> data = new Dictionary<char, int>{
        {'I', 1},
        {'V', 5},
        {'X', 10},
        {'L', 50},
        {'C', 100},
        {'D', 500},
        {'M', 1000}
    };

    public int RomanToInt(string s)
    {
        int length = s.Length;
        var source = s.ToCharArray();

        int result = 0;
        string unit = "";

        for (int i = 0; i < length; i++)
        {
            unit += source[i];

            if (i < length - 1 && data[source[i]] < data[source[i + 1]])
            {
                continue;
            }

            // 取出unit的值
            var temp = unit.ToCharArray();
            if (temp.Length == 1)
            {
                result += data[temp[0]];
            }
            else
            {
                for (int j = 0; j < temp.Length; j++)
                {
                    if (j != temp.Length - 1)
                    {
                        result -= data[temp[j]];
                    }
                    else
                    {
                        result += data[temp[j]];
                    }
                }
            }

            unit = "";
        }

        return result;
    }
}
// @lc code=end

