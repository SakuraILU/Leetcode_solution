/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	xors := 0
	for _, nums := range nums {
		xors ^= nums
	}

	partition := xors & (-xors)

	res := []int{0, 0}
	for _, num := range nums {
		if (num & partition) != 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}

	return res
}

// @lc code=end

