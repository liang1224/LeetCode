using NUnit.Framework;

namespace LeetCodeTest
{
    public class Tests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void Test1()
        {
            var a = "aaaaa";
            var b = "bba";

            var result = StrStr(a, b);

            Assert.AreEqual(2, result);
        }

        public int StrStr(string haystack, string needle)
        {
            if (needle.Length == 0)
            {
                return 0;
            }

            if (haystack.Length == 0 || needle.Length > haystack.Length)
            {
                return -1;
            }

            var next = BuildNext(needle);

            int s = 0, c = 0;

            while (s < haystack.Length)
            {
                if (haystack[s] == needle[c])
                {
                    c++;

                    if (c >= needle.Length)
                    {
                        return s - c;
                    }
                }
                else
                {
                    if (s == 0)
                    {
                        c = 0;
                    }
                    else
                    {
                        c = next[c - 1];
                    }
                }

                s++;
            }

            return -1;
        }

        private int[] BuildNext(string p)
        {
            int[] next = new int[p.Length];
            next[0] = 0;

            for (int slow = 0, fast = 1; fast < p.Length && slow < fast; fast++)
            {
                if (p[fast] == p[slow])
                {
                    next[fast] = next[fast - 1] + 1;
                    slow++;
                }
                else
                {
                    if (slow != 0)
                    {
                        slow = next[slow - 1];
                    }
                    else
                    {
                        next[fast] = 0;
                    }
                }
            }

            return next;
        }
    }
}