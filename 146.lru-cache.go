/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type Entry struct {
	Key int
	Val int
}
type LRUCache struct {
	kvs      map[int]*list.Element
	lst      *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kvs:      make(map[int]*list.Element),
		lst:      list.New().Init(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.kvs[key]
	if !ok {
		return -1
	}
	entry := elem.Value.(*Entry)

	this.lst.MoveToBack(elem)

	return entry.Val
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.kvs[key]; ok {
		entry := elem.Value.(*Entry)
		entry.Val = value
		this.lst.MoveToBack(elem)
		return
	}

	if len(this.kvs) >= this.capacity {
		this.evict()
	}

	entry := &Entry{Key: key, Val: value}
	elem := this.lst.PushBack(entry)
	this.kvs[key] = elem
}

func (this *LRUCache) evict() {
	elem := this.lst.Front()
	entry := elem.Value.(*Entry)
	delete(this.kvs, entry.Key)
	this.lst.Remove(elem)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
