/*
 * @lc app=leetcode id=1541 lang=golang
 *
 * [1541] Minimum Insertions to Balance a Parentheses String
 */

// @lc code=start
import "fmt"

type Queue struct {
	elems []rune
}

func NewQueue() *Queue {
	return &Queue{
		elems: make([]rune, 0),
	}
}

func (q *Queue) Push(val rune) {
	q.elems = append(q.elems, val)
}

func (q *Queue) Pop() (elem rune, err error) {
	if len(q.elems) == 0 {
		err = fmt.Errorf("empty")
		return
	}
	elem = q.elems[len(q.elems)-1]
	q.elems = q.elems[0 : len(q.elems)-1]
	return
}

func (q *Queue) Size() int {
	return len(q.elems)
}

func minInsertions(s string) int {
	queue := NewQueue()
	need := 0
	for _, elem := range s {
		switch elem {
		case '(':
			if queue.Size()%2 != 0 {
				need++
				queue.Pop()
			}
			queue.Push(elem)
			queue.Push(elem)
		case ')':
			_, err := queue.Pop()
			if err != nil {
				need++
				queue.Push('(')
			}
		}
	}

	need += queue.Size()
	return need
}

// @lc code=end
