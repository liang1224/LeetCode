/*
 * @lc app=leetcode.cn id=20 lang=csharp
 *
 * [20] 有效的括号
 */

// @lc code=start
public class Solution
{
    public bool IsValid(string s)
    {
        Stack<char> stack = new Stack<char>();

        foreach (var c in s)
        {
            if (c == '(' || c == '[' || c == '{')
            {
                stack.Push(c);
            }
            else
            {
                if (stack.Count == 0)
                {
                    return false;
                }

                switch (c)
                {
                    case ')':
                        if (stack.Pop() != '(')
                        {
                            return false;
                        }
                        break;
                    case ']':
                        if (stack.Pop() != '[')
                        {
                            return false;
                        }
                        break;
                    case '}':
                        if (stack.Pop() != '{')
                        {
                            return false;
                        }
                        break;
                    default:
                        return false;
                }
            }
        }

        return stack.Count == 0;
    }
}
// @lc code=end

