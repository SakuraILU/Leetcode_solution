/*
 * @lc app=leetcode id=845 lang=golang
 *
 * [845] Longest Mountain in Array
 */

// @lc code=start
var _arr []int
var memo map[int]int

func longestMountain(arr []int) int {
	premlen := longestIncrease(arr)
	aftermlen := longestDecrease(arr)

	mlen := 0
	for i, _ := range premlen {
		if premlen[i] > 1 && aftermlen[i] > 1 {
			mlen = max(mlen, premlen[i]+aftermlen[i]-1)
		}
	}

	return mlen
}

func longestIncrease(arr []int) (res []int) {
	_arr = arr
	memo = make(map[int]int)

	res = make([]int, len(_arr))
	for i := 0; i < len(_arr); i++ {
		res[i] = dp(i)
	}
	return
}

func dp(i int) (mnum int) {
	if i == 0 {
		return 1
	}

	if v, ok := memo[i]; ok {
		return v
	}

	mnum = 1
	if _arr[i] > _arr[i-1] {
		mnum += dp(i - 1)
	}

	memo[i] = mnum
	return
}

func longestDecrease(arr []int) (res []int) {
	reverse(arr)
	res = longestIncrease(arr)
	reverse(res)
	return
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp

		left++
		right--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
