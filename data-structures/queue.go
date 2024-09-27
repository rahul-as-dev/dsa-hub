package data_structures

import (
	"fmt"
	"math/rand/v2"
)

/*
*
 */
func Main5() {
	var q []int
	q = queueExample(q)
	/*allocates an underlying array
	of size 10 and returns a slice of length 0 and capacity 10
	returns a slice header as struct {
		p &int
		len int
		cap int
	}*/
	q1 := make([]int, 0, 5)
	// IMP - header slices are passed by value, so if you do any reassign, it is not reflected in the original slice header
	// in the caller. However, if you just do lets say - q[1] = new_value, it will be reflected in the caller as the underlying
	// array is same and in that case, header value remains unchanged.
	q1 = queueExample(q1)

	fmt.Println(len(q))
}

/*
*
A normal FIFO structure. Supports enqueue and peek in O(1) time and polling in O(n) worst case.
Slice is used which is basically a pointer to underlying array (queue)
*/
func queueExample(q []int) []int {
	//Enqueuing (offer)
	for range 3 {
		q = append(q, rand.IntN(100))
	}
	fmt.Println(q)
	// Peeking
	peek := q[0]
	fmt.Println(peek)
	//Polling
	//O(n) in the worst case. This is because slicing off the front of the slice involves creating a new slice header and
	//can involve reallocation of the underlying array if the slice is not implemented as a circular buffer.
	q = q[1:]
	fmt.Println(q)
	q = append(q, rand.IntN(100))
	fmt.Println(q)
	return q
}

// CircularQueue When you already have idea about capacity
type CircularQueue struct {
	data  []int
	size  int // capacity of the queue
	front int // index of the front element
	rear  int // index of the rear element
	count int // current number of elements
}

// Create a new CircularQueue with the given capacity
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, capacity),
		size:  capacity,
		front: 0,
		rear:  0,
		count: 0,
	}
}

// Enqueue adds an element to the rear of the queue
func (q *CircularQueue) Enqueue(value int) bool {
	if q.count == q.size {
		return false // Queue is full
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	q.count++
	return true
}

// Dequeue removes and returns the front element of the queue
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.count == 0 {
		return 0, false // Queue is empty
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % q.size
	q.count--
	return value, true
}

// Peek returns the front element without removing it
func (q *CircularQueue) Peek() (int, bool) {
	if q.count == 0 {
		return 0, false // Queue is empty
	}
	return q.data[q.front], true
}
