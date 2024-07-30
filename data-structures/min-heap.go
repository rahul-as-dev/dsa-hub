package main

import "fmt"

type HeapNode struct {
	score1, score2 int
}
type PriorityQueue struct {
	data []HeapNode
	cap  int
}

func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{data: make([]HeapNode, 0), cap: cap}
}

// Enqueue (We have to retain the property of heap for each insertion/deletion)
func (pq *PriorityQueue) Enqueue(hn HeapNode) bool {
	if len(pq.data) == pq.cap {
		return false
	}
	// first add the element
	pq.data = append(pq.data, HeapNode{hn.score1, hn.score2})
	// then repeatedly swap it with its parent till the condition - element >= parent
	length := len(pq.data)
	parentInd := (length / 2) - 1
	hnIndex := length - 1

	for parentInd >= 0 && hn.score1+hn.score2 < pq.data[parentInd].score1+pq.data[parentInd].score2 {
		pq.data[parentInd], pq.data[hnIndex] = pq.data[hnIndex], pq.data[parentInd]
		hnIndex = parentInd
		parentInd = ((hnIndex + 1) / 2) - 1
	}
	return true
}

// Create a min heap (PriorityQueue) of HeapNode with priority of score1+score2
func main() {
	/*A heap is just a complete binary tree.
	A min heap have its each node's value lesser than its decendent
	*/
	var pq []int
	/*A heap can be visualized and implemented from a slice/array with below constraint
		1. Traverse the binary tree layer wise from top (left-to-right)
		2. Maintain an array (1-based index) of those node satisfying the property - left-child @ 2*i index; right child @ (2*i + 1) index;
			and parent @ floor(i/2) index.
	Inserting/Deleting can be done in O(log(N)) time
	*/
	//pq = PriorityQueueExample(pq)
	fmt.Println(pq)
}
