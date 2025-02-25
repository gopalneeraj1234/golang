package recentcall933

type Queue struct {
	elements []int
}

func (q Queue) isEmpty() bool {
	return q.size() == 0
}

func (q Queue) size() int {
	return len(q.elements)
}

func (q *Queue) add(time int) {
	q.elements = append(q.elements, time)
}

func (q *Queue) head() int {
	return q.elements[0]
}

func (q *Queue) remove() int {
	retVal := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return retVal
}

func NewQueue() *Queue {
	q := &Queue{
		elements: make([]int, 0),
	}
	return q
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
