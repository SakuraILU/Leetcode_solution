/*
 * @lc app=leetcode id=639 lang=golang
 *
 * [639] Decode Ways II
 */

// @lc code=start
var memo map[int]int

var _s string

// Since the answer may be very large, return it modulo 10**9 + 7.
var M = int(math.Pow(10, 9)) + 7

func numDecodings(s string) (cnt int) {
	_s = s

	memo = make(map[int]int)
	return dp(0)
}

func dp(start int) (cnt int) {
	if start >= len(_s) {
		return 1
	}

	if v, ok := memo[start]; ok {
		return v
	}

	for i := start; i < min(start+2, len(_s)); i++ {
		snum := _s[start : i+1]
		if strings.HasPrefix(snum, "0") {
			continue
		}

		k := decodeways(snum)
		if k == 0 {
			continue
		}

		cnt = (cnt + k*dp(i+1)) % M // don't forget to mod cnt by M
	}

	memo[start] = cnt
	return
}

func decodeways(s string) int {
	// note: * can stand for 1 to 9 (0 is excluded)
	if len(s) == 1 {
		if s[0] == '*' {
			return 9
		} else {
			return 1
		}
	}

	if len(s) == 2 {
		if s[0] == '*' {
			if s[1] == '*' {
				return 15 // 11-19, 21-26
			} else if s[1] >= '0' && s[1] <= '6' {
				return 2
			} else {
				return 1
			}
		} else if s == "1*" {
			return 9 // 11-19
		} else if s == "2*" {
			return 6 // 21-26
		} else {
			num, _ := strconv.Atoi(s)
			if num > 0 && num <= 26 {
				return 1
			}
		}
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
