/*
 * @lc app=leetcode id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 */

// @lc code=start
var memolmax map[int]int
var memormin map[int]int

var _nums []int

func maxSubarraySumCircular(nums []int) int {
	_nums = nums
	memolmax = make(map[int]int)
	memormin = make(map[int]int)

	totalsum := 0
	for _, num := range nums {
		totalsum += num
	}

	msum := math.MinInt
	for i := 0; i < len(nums); i++ {
		lmax := dplmax(i)
		rmin := 0
		if i < len(nums)-1 {
			rmin = dprmin(i + 1)
		}
		sum := max(lmax, totalsum-rmin)
		msum = max(msum, sum)
	}

	return msum
}

func dplmax(i int) (msum int) {
	if i == 0 {
		return _nums[i]
	}

	if v, ok := memolmax[i]; ok {
		return v
	}

	msum = max(_nums[i]+dplmax(i-1), _nums[i])

	memolmax[i] = msum
	return msum
}

func dprmin(i int) (msum int) {
	if i == len(_nums)-1 {
		return _nums[i]
	}

	if v, ok := memormin[i]; ok {
		return v
	}

	msum = min(_nums[i]+dprmin(i+1), _nums[i])

	memormin[i] = msum
	return msum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
