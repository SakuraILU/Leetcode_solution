/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */

// @lc code=start
type IntHeap struct {
	ints []int
}

func NewIntHeap() *IntHeap {
	return &IntHeap{
		ints: make([]int, 0),
	}
}

func (ih *IntHeap) Push(v interface{}) {
	ih.ints = append(ih.ints, v.(int))
}

func (ih *IntHeap) Less(i, j int) bool {
	return ih.ints[i] > ih.ints[j]
}

func (ih *IntHeap) Swap(i, j int) {
	tmp := ih.ints[i]
	ih.ints[i] = ih.ints[j]
	ih.ints[j] = tmp
}

func (ih *IntHeap) Pop() interface{} {
	v := ih.ints[len(ih.ints)-1]
	ih.ints = ih.ints[0 : len(ih.ints)-1]
	return v
}

func (ih *IntHeap) Len() int {
	return len(ih.ints)
}

func (ih *IntHeap) Top() int {
	if len(ih.ints) == 0 {
		return 0
	}

	return ih.ints[0]
}

func lastStoneWeight(stones []int) int {
	intheap := NewIntHeap()
	for _, stone := range stones {
		heap.Push(intheap, stone)
	}
	heap.Init(intheap)

	for intheap.Len() > 1 {
		y := heap.Pop(intheap).(int)
		x := heap.Pop(intheap).(int)

		if x != y {
			heap.Push(intheap, y-x)
		}
	}

	return intheap.Top()
}

// @lc code=end
