/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
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

var match map[rune]rune = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	queue := NewQueue()

	for _, elem := range s {
		switch elem {
		case '(', '[', '{':
			queue.Push(elem)
		case ')', ']', '}':
			val, err := queue.Pop()
			if err != nil || match[val] != elem {
				return false
			}
		}
	}

	return queue.Size() == 0
}

// @lc code=end
