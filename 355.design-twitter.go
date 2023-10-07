/*
 * @lc app=leetcode id=355 lang=golang
 *
 * [355] Design Twitter
 */

// @lc code=start

type MsgHeap struct {
	msgs []*list.Element
}

func NewIntsHeap() *MsgHeap {
	return &MsgHeap{
		msgs: make([]*list.Element, 0),
	}
}

func (ih *MsgHeap) Push(v interface{}) {
	ih.msgs = append(ih.msgs, v.(*list.Element))
}

func (ih *MsgHeap) Less(i, j int) bool {
	msg1 := ih.msgs[i].Value.(*TwitterMsg)
	msg2 := ih.msgs[j].Value.(*TwitterMsg)
	return msg1.time.After(msg2.time)
}

func (ih *MsgHeap) Swap(i, j int) {
	tmp := ih.msgs[i]
	ih.msgs[i] = ih.msgs[j]
	ih.msgs[j] = tmp
}

func (ih *MsgHeap) Pop() interface{} {
	v := ih.msgs[len(ih.msgs)-1]
	ih.msgs = ih.msgs[0 : len(ih.msgs)-1]
	return v
}

func (ih *MsgHeap) Len() int {
	return len(ih.msgs)
}

type TwitterMsg struct {
	tweetId int
	time    time.Time
}

type User struct {
	userId      int
	twitterlist *list.List
}

type Twitter struct {
	users       map[int]User
	followlists map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		users:       make(map[int]User),
		followlists: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.LazyInitUser(userId)

	msg := &TwitterMsg{
		tweetId: tweetId,
		time:    time.Now(),
	}

	user := this.users[userId]
	user.twitterlist.PushFront(msg)

	return
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.LazyInitUser(userId)

	msgheap := NewIntsHeap()
	heap.Init(msgheap)
	if own_msg := this.users[userId].twitterlist.Front(); own_msg != nil {
		heap.Push(msgheap, own_msg)
	}

	userfollows := this.followlists[userId]
	for followeeId, _ := range userfollows {
		if msg := this.users[followeeId].twitterlist.Front(); msg != nil {
			heap.Push(msgheap, msg)
		}
	}

	recent10 := make([]int, 0)
	for msgheap.Len() > 0 {
		elem := heap.Pop(msgheap).(*list.Element)
		msg := elem.Value.(*TwitterMsg)
		recent10 = append(recent10, msg.tweetId)

		if len(recent10) >= 10 {
			break
		}

		if nelem := elem.Next(); nelem != nil {
			heap.Push(msgheap, elem.Next())
		}
	}

	return recent10
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.LazyInitUser(followerId)
	this.LazyInitUser(followeeId)

	userfollows := this.followlists[followerId]
	userfollows[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.LazyInitUser(followerId)
	this.LazyInitUser(followeeId)

	userfollows := this.followlists[followerId]
	delete(userfollows, followeeId)
}

func (this *Twitter) LazyInitUser(userId int) {
	if _, ok := this.users[userId]; !ok {
		this.users[userId] = User{
			userId:      userId,
			twitterlist: list.New().Init(),
		}
	}
	if _, ok := this.followlists[userId]; !ok {
		this.followlists[userId] = make(map[int]bool)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end
