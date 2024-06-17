package leetcode

type Queue[T any] struct {
	s []T
	h int
}

func (q *Queue[T]) Enqueue(e T) {
	q.s = append(q.s, e)
	q.h++
}

func (q *Queue[T]) Dequeue() T {
	if q.h < 1 {
		var e T
		return e
	}

	e := q.s[0]
	q.s = q.s[1:]
	q.h--
	return e
}

func (q *Queue[T]) Peek() T {
	e := q.s[0]
	return e
}

func (q *Queue[T]) Height() int {
	return q.h
}

type RecentCounter struct {
	requests Queue[int]
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: Queue[int]{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests.Enqueue(t)

	for this.requests.Peek() < t-3000 {
		this.requests.Dequeue()
	}

	return this.requests.Height()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

// # Description:
//
// You have a RecentCounter class which counts the number of recent requests within a certain time
// frame.
//
// Implement the RecentCounter class:
//
// - `RecentCounter()` Initializes the counter with zero recent requests.
// - `int ping(int t)` Adds a new request at time t, where t represents some time in milliseconds,
// and returns the number of requests that has happened in the past 3000 milliseconds (including the
// new request). Specifically, return the number of requests that have happened in the inclusive
// range [t - 3000, t].
// It is guaranteed that every call to ping uses a strictly larger value of t than the previous
// call.
//
// https://leetcode.com/problems/number-of-recent-calls
//
// Notes: Use a queue and dequeue FIFO as soon as we find a response time out of bounds since we are
// told all subsequent requests to Ping will be larger
//
// tags: [Queue, easy]
