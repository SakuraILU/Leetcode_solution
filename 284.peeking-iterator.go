/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
const EMPTY = 0

type PeekingIterator struct {
	iter    *Iterator
	peekbuf int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:    iter,
		peekbuf: 0,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.peekbuf != EMPTY || this.iter.hasNext()
}

func (this *PeekingIterator) next() (val int) {
	if this.peekbuf != EMPTY {
		val = this.peekbuf
		this.peekbuf = EMPTY
	} else {
		val = this.iter.next()
	}

	return val
}

func (this *PeekingIterator) peek() int {
	if this.peekbuf != EMPTY {
		return this.peekbuf
	}

	this.peekbuf = this.iter.next()
	return this.peekbuf
}

// @lc code=end
