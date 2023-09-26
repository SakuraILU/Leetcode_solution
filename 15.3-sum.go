/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
var _nums []int

var track []int
var cursum int
var res [][]int

const (
	LEFT = iota
	RIGHT
)

func threeSum(nums []int) [][]int {
	_nums = nums
	sort.Ints(_nums)
	// fmt.Println(_nums)

	track = make([]int, 0)
	cursum = 0
	res = make([][]int, 0)

	dfs(0, 3)

	return res
}

func dfs(start int, n int) {
	if n == 2 {
		if len(_nums)-start < 2 {
			return
		}
		twoSum(start, len(_nums)-1, -cursum)
		return
	}

	for i := start; i < len(_nums); i++ {
		if i > start && _nums[i] == _nums[i-1] {
			continue
		}

		track = append(track, _nums[i])
		cursum += _nums[i]

		dfs(i+1, n-1)

		cursum -= _nums[i]
		track = track[0 : len(track)-1]
	}
}

func twoSum(left, right int, target int) {
	for left < right {
		sum := _nums[left] + _nums[right]
		if sum > target {
			right = skipReplica(right, LEFT)
		} else if sum < target {
			left = skipReplica(left, RIGHT)
		} else {
			track = append(track, _nums[left], _nums[right])

			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)

			track = track[0 : len(track)-2]

			left = skipReplica(left, RIGHT)
			right = skipReplica(right, LEFT)
		}
	}
}

func skipReplica(pos, dir int) (next int) {
	if dir == LEFT {
		next = pos - 1
		for next > 0 && _nums[next] == _nums[pos] {
			next--
		}
	} else {
		next = pos + 1
		for next < len(_nums) && _nums[next] == _nums[pos] {
			next++
		}
	}
	return
}

// @lc code=end
