package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	priority int
}

// Note that you have to implement a total of 5 methods to use the inbuilt heap
// TODO- Look for Generic use to reduce the above painpoint
//The downsides to the interface approach, besides having to implement five functions, are the interface type assertions from
//pop and push. The caller of pop needs to type assert, which is only a few lines of code.
//But passing the wrong type to push could lead to a panic, as it does in this example.

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int      { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq *PriorityQueue) Push(x interface{}) {
	// a failed type assertion here will panic
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := &PriorityQueue{}
	heap.Push(pq, &Item{value: 100, priority: 10})

	item := heap.Pop(pq)
	fmt.Println(item)
}
