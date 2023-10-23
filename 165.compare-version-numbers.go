/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start

func compareVersion(version1 string, version2 string) int {
	snums1 := strings.Split(version1, ".")
	snums2 := strings.Split(version2, ".")

	i := 0
	for ; i < len(snums1) && i < len(snums2); i++ {
		num1, _ := strconv.Atoi(snums1[i])
		num2, _ := strconv.Atoi(snums2[i])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	if len(snums1) > len(snums2) {
		for ; i < len(snums1); i++ {
			if n, _ := strconv.Atoi(snums1[i]); n != 0 {
				return 1
			}
		}
	} else if len(snums1) < len(snums2) {
		for ; i < len(snums2); i++ {
			if n, _ := strconv.Atoi(snums2[i]); n != 0 {
				return -1
			}
		}
	}
	return 0
}

// @lc code=end
