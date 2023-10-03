/*
 * @lc app=leetcode id=1845 lang=golang
 *
 * [1845] Seat Reservation Manager
 */

// @lc code=start

type IntsHeap struct {
	ints []int
}

func NewIntsHeap() *IntsHeap {
	return &IntsHeap{
		ints: make([]int, 0),
	}
}

func (ih *IntsHeap) Push(v interface{}) {
	ih.ints = append(ih.ints, v.(int))
}

func (ih *IntsHeap) Less(i, j int) bool {
	return ih.ints[i] < ih.ints[j]
}

func (ih *IntsHeap) Swap(i, j int) {
	tmp := ih.ints[i]
	ih.ints[i] = ih.ints[j]
	ih.ints[j] = tmp
}

func (ih *IntsHeap) Pop() interface{} {
	v := ih.ints[len(ih.ints)-1]

	ih.ints = ih.ints[0 : len(ih.ints)-1]

	return v
}

func (ih *IntsHeap) Len() int {
	return len(ih.ints)
}

func (ih *IntsHeap) Top() int {
	return ih.ints[0]
}

type SeatManager struct {
	intheap *IntsHeap
}

func Constructor(n int) SeatManager {
	sh := SeatManager{
		intheap: NewIntsHeap(),
	}

	for i := 1; i <= n; i++ {
		heap.Push(sh.intheap, i)
	}

	return sh
}

func (this *SeatManager) Reserve() int {
	top := heap.Pop(this.intheap).(int)
	return top
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.intheap, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
