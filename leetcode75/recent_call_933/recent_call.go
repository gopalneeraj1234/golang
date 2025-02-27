package recentcall933

type Queue []int

func (q Queue) isEmpty() bool {
	return q.size() == 0
}

func (q Queue) size() int {
	return len(q)
}

func (q *Queue) add(time int) {
	*q = append(*q, time)
}

func (q *Queue) head() int {
	return (*q)[0]
}

func (q *Queue) remove() int {
	retVal := (*q)[0]
	*q = (*q)[1:len(*q)]
	return retVal
}

func NewQueue() *Queue {
	q := make(Queue, 0)
	return &q
}

type RecentCounter struct {
	q Queue
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: *NewQueue(),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.q.add(t)
	for !this.q.isEmpty() && this.q.head() < t-3000 {
		this.q.remove()
	}
	return this.q.size()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
