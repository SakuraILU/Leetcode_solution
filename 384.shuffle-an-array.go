/*
 * @lc app=leetcode id=384 lang=golang
 *
 * [384] Shuffle an Array
 */

// @lc code=start
type Solution struct {
	orgnums []int
	nums    []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().Unix())
	s := Solution{
		orgnums: make([]int, len(nums)),
		nums:    nums,
	}

	copy(s.orgnums, s.nums)

	return s
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.orgnums)
	return this.orgnums
}

func (this *Solution) Shuffle() []int {
	nums := this.nums
	for i := 0; i < len(nums); i++ {
		j := i + rand.Intn(len(nums)-i)

		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}

	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
