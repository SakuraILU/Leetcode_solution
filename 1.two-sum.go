/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) (res []int) {
	table := make(map[int]int, 0)
	for idx, num := range nums {
		table[num] = idx
	}

	for idx, num := range nums {
		if v, ok := table[target-num]; ok {
			if idx != v {
				res = append(res, idx, v)
				return
			}
		}
	}
	return
}

// @lc code=end
