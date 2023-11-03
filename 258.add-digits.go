/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
	for num >= 10 {
		num = next(num)
	}

	return num
}

func next(num int) (res int) {
	for num > 0 {
		digit := num % 10
		res += digit
		num /= 10
	}
	return
}

// @lc code=end

