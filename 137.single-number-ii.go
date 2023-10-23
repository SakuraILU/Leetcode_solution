/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			if (num & (1 << i)) != 0 {
				cnt++
			}
		}
		if cnt%3 != 0 {
			res |= (1 << i)
		}
	}

	return int(res)
}

// @lc code=end
