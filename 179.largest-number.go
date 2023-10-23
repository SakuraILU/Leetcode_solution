/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) string {
	snums := make([]string, 0)
	for _, num := range nums {
		snums = append(snums, strconv.Itoa(num))
	}

	sort.Slice(snums, func(i, j int) bool {
		return snums[i]+snums[j] > snums[j]+snums[i]
	})

	if snums[0] == "0" {
		return "0"
	}
	return strings.Join(snums, "")
}

// @lc code=end
