/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	fast, slow := n, n

	for {
		fast = next(next(fast))
		slow = next(slow)

		if slow == fast {
			break
		}
	}

	return slow == 1
}

func next(n int) (num int) {
	for n > 0 {
		digit := n % 10
		n = n / 10
		num += digit * digit
	}
	return
}

// @lc code=end

