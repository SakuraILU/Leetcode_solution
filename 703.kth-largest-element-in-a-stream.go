/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
type IntsHeap struct {
	arr []int
}

func NewIntsHeap() *IntsHeap {
	return &IntsHeap{
		arr: make([]int, 0),
	}
}

func (ih *IntsHeap) Push(v interface{}) {
	ih.arr = append(ih.arr, v.(int))
}

func (ih *IntsHeap) Pop() interface{} {
	v := ih.arr[len(ih.arr)-1]
	ih.arr = ih.arr[0 : len(ih.arr)-1]

	return v
}

func (ih *IntsHeap) Less(i, j int) bool {
	return ih.arr[i] < ih.arr[j]
}

func (ih *IntsHeap) Swap(i, j int) {
	tmp := ih.arr[i]
	ih.arr[i] = ih.arr[j]
	ih.arr[j] = tmp

	return
}

func (ih *IntsHeap) Len() int {
	return len(ih.arr)
}

func (ih *IntsHeap) Top() int {
	return ih.arr[0]
}

type KthLargest struct {
	intsheap *IntsHeap
	k        int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		intsheap: NewIntsHeap(),
		k:        k,
	}

	sort.Ints(nums)
	heap.Init(kl.intsheap)

	for i := max(len(nums)-k, 0); i < len(nums); i++ {
		heap.Push(kl.intsheap, nums[i])
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.intsheap.Len() < this.k {
		heap.Push(this.intsheap, val)
		return this.intsheap.Top()
	}

	kth := this.intsheap.Top()
	if val > kth {
		heap.Pop(this.intsheap)
		heap.Push(this.intsheap, val)
	}

	return this.intsheap.Top()
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
