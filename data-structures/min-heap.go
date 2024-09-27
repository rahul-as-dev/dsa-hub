package data_structures

import "fmt"

type HeapNode struct {
	score1, score2 int
}
type CustomPriorityQueue struct {
	data   []HeapNode
	length int
	cap    int
}

func NewCustomPriorityQueue(cap int) *CustomPriorityQueue {
	return &CustomPriorityQueue{data: make([]HeapNode, 0), length: 0, cap: cap}
}

// Enqueue (We have to retain the property of heap for each insertion/deletion)
func (pq *CustomPriorityQueue) Enqueue(hn HeapNode) bool {
	if pq.length == pq.cap {
		return false
	}
	// first add the element
	pq.data = append(pq.data, HeapNode{hn.score1, hn.score2})
	pq.length++
	// then repeatedly swap it with its parent till the condition - element >= parent
	parentInd := (pq.length / 2) - 1
	hnIndex := pq.length - 1

	for parentInd >= 0 && hn.score1+hn.score2 < pq.data[parentInd].score1+pq.data[parentInd].score2 {
		pq.data[parentInd], pq.data[hnIndex] = pq.data[hnIndex], pq.data[parentInd]
		hnIndex = parentInd
		parentInd = ((hnIndex + 1) / 2) - 1
	}
	return true
}

// Dequeue (We have to retain the property of heap for each insertion/deletion)
func (pq *CustomPriorityQueue) Dequeue() (HeapNode, bool) {
	if pq.length == 0 {
		return HeapNode{}, false
	}

	res := pq.data[0]
	hn := pq.data[pq.length-1]
	pq.data[0] = pq.data[pq.length-1]
	pq.data = pq.data[:pq.length-1]
	pq.length--
	hnIndex := 0
	var swapIndex int
	for hnIndex < pq.length {
		leftInd, rightInd := 2*hnIndex+1, 2*hnIndex+2
		if leftInd < pq.length && rightInd < pq.length {
			if hn.score1+hn.score2 > min(pq.data[leftInd].score1+pq.data[leftInd].score1, pq.data[rightInd].score1+pq.data[rightInd].score2) {
				if pq.data[leftInd].score1+pq.data[leftInd].score2 < pq.data[rightInd].score1+pq.data[rightInd].score2 {
					swapIndex = leftInd
				} else {
					swapIndex = rightInd
				}
				pq.data[hnIndex], pq.data[swapIndex] = pq.data[swapIndex], pq.data[hnIndex]
				hnIndex = swapIndex
			} else {
				break
			}
		} else if leftInd < pq.length && rightInd >= pq.length {
			if hn.score1+hn.score2 > pq.data[leftInd].score1+pq.data[leftInd].score2 {
				swapIndex = leftInd
				pq.data[hnIndex], pq.data[swapIndex] = pq.data[swapIndex], pq.data[hnIndex]
				hnIndex = swapIndex
			} else {
				break
			}
		} else {
			break
		}
	}
	return res, true
}

// Create a min heap (CustomPriorityQueue) of HeapNode with priority of score1+score2
func Main3() {
	/*A heap is just a complete binary tree.
	A min heap have its each node's value lesser than its decendent
	*/
	/*A heap can be visualized and implemented from a slice/array with below constraint
		1. Traverse the binary tree layer wise from top (left-to-right)
		2. Maintain an array (1-based index) of those node satisfying the property - left-child @ 2*i index; right child @ (2*i + 1) index;
			and parent @ floor(i/2) index.
	Inserting/Deleting can be done in O(log(N)) time
	*/
	pq := NewCustomPriorityQueue(5)
	pq.Enqueue(HeapNode{1, 0})
	fmt.Println(pq.data)
	pq.Enqueue(HeapNode{2, 1})
	fmt.Println(pq.data)
	pq.Enqueue(HeapNode{1, 1})
	fmt.Println(pq.data)
	pq.Enqueue(HeapNode{-1, 0})
	pq.Enqueue(HeapNode{-1, -1})
	fmt.Println(pq)
	el, _ := pq.Dequeue()
	fmt.Println(pq)
	fmt.Println(el)
	el1, _ := pq.Dequeue()
	fmt.Println(pq)
	fmt.Println(el1)
	el2, _ := pq.Dequeue()
	fmt.Println(pq)
	fmt.Println(el2)
	el3, _ := pq.Dequeue()
	fmt.Println(pq)
	fmt.Println(el3)
	el4, _ := pq.Dequeue()
	fmt.Println(pq)
	fmt.Println(el4)
	pq.Enqueue(HeapNode{-10, -1})
	fmt.Println(pq)
}
