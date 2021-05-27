/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	base = make(map[byte]int)
	base["I"[0]] = 1
	base["V"[0]] = 5
	base["X"[0]] = 10
	base["L"[0]] = 50
	base["C"[0]] = 100
	base["D"[0]] = 500
	base["L"[0]] = 1000

	len := len(s)
	result := 0

	for i, v := range s {
		unit := string(v)

		if int(i+1) < len {
			j := i

			for base[s[j]] < base[s[j+1]] || (base[s[j]] == base[s[j+1]] && s[j] == "I"[0]) {
				unit = unit + string(s[j+1])
			}
		}

		result = transPart(unit, result)
	}

	return result
}

var base map[byte]int

func transPart(s string, origin int) int {
	if len(s) == 1 {
		return origin*10 + base[s[0]]
	} else if s[len(s)-1] == "I"[0] {
		return origin + 10 + len(s)
	} else {
		return origin*10 + (base[s[1]] - base[s[0]])
	}
}

// @lc code=end

