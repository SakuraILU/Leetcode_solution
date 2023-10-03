/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
type Entry struct {
	word string
	freq int
}

type EntryHeap struct {
	ints []*Entry
}

func NewEntryHeap() *EntryHeap {
	return &EntryHeap{
		ints: make([]*Entry, 0),
	}
}

func (ih *EntryHeap) Push(v interface{}) {
	ih.ints = append(ih.ints, v.(*Entry))
}

func (ih *EntryHeap) Less(i, j int) bool {
	entrya := ih.ints[i]
	entryb := ih.ints[j]
	if entrya.freq > entryb.freq {
		return true
	} else if entrya.freq == entryb.freq && entrya.word < entryb.word {
		return true
	}

	return false
}

func (ih *EntryHeap) Swap(i, j int) {
	tmp := ih.ints[i]
	ih.ints[i] = ih.ints[j]
	ih.ints[j] = tmp
}

func (ih *EntryHeap) Pop() interface{} {
	v := ih.ints[len(ih.ints)-1]
	ih.ints = ih.ints[0 : len(ih.ints)-1]
	return v
}

func (ih *EntryHeap) Len() int {
	return len(ih.ints)
}

func (ih *EntryHeap) Top() *Entry {
	if len(ih.ints) == 0 {
		return nil
	}

	return ih.ints[0]
}

func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int, 0)
	for _, word := range words {
		if _, ok := counter[word]; !ok {
			counter[word] = 1
		} else {
			counter[word]++
		}
	}

	entryheap := NewEntryHeap()
	for word, freq := range counter {
		fmt.Println(word, freq)
		entry := &Entry{
			word: word,
			freq: freq,
		}

		entryheap.Push(entry)
	}
	heap.Init(entryheap)

	topk := make([]string, 0)
	for i := 0; i < k; i++ {
		val := heap.Pop(entryheap).(*Entry)
		topk = append(topk, val.word)
	}

	return topk
}

// @lc code=end
