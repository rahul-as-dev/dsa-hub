package data_structures

import "fmt"

// Not handling the capacity as it is not required for most dsa problems
type GenericPriorityQueue[T any] struct {
	data    []T
	compare func(t1, t2 T) bool
	length  int
}

func NewGenericPriorityQueue[T any](compare func(t1, t2 T) bool) *GenericPriorityQueue[T] {
	return &GenericPriorityQueue[T]{data: make([]T, 0), compare: compare, length: 0}
}

func (gpq *GenericPriorityQueue[T]) Len() int {
	return gpq.length
}
func (gpq *GenericPriorityQueue[T]) Push(t T) {
	gpq.data = append(gpq.data, t)
	gpq.length++
	gpq.heapifyUpwards(gpq.length - 1)
}

func (gpq *GenericPriorityQueue[T]) heapifyUpwards(i int) {
	for i > 0 && gpq.compare(gpq.data[i], gpq.data[(i-1)/2]) {
		gpq.data[i], gpq.data[(i-1)/2] = gpq.data[(i-1)/2], gpq.data[i]
		i = (i - 1) / 2
	}
}

func (gpq *GenericPriorityQueue[T]) Pop() T {
	var value T
	if gpq.length == 0 {
		return value
	}
	value = gpq.data[0]
	gpq.data[0] = gpq.data[gpq.length-1]
	//gpq.data[gpq.length - 1] = nil
	gpq.data = gpq.data[:gpq.length-1]
	gpq.length--
	gpq.heapifyDownwards(0)
	return value
}

func (gpq *GenericPriorityQueue[T]) heapifyDownwards(i int) {
	leftInd, rightInd := 2*i+1, 2*i+2
	swapedInd := i
	if leftInd < gpq.length && !gpq.compare(gpq.data[i], gpq.data[leftInd]) {
		gpq.data[i], gpq.data[leftInd] = gpq.data[leftInd], gpq.data[i]
		swapedInd = leftInd
	}
	if rightInd < gpq.length && !gpq.compare(gpq.data[i], gpq.data[rightInd]) {
		gpq.data[i], gpq.data[rightInd] = gpq.data[rightInd], gpq.data[i]
		swapedInd = rightInd
	}
	if swapedInd != i {
		i = swapedInd
		gpq.heapifyDownwards(i)
	}
}

//	func (gpq *GenericPriorityQueue[T]) heapifyDownwards(i int) {
//		for {
//			leftInd, rightInd := 2*i+1, 2*i+2
//			largest := i
//
//			if leftInd < gpq.length && gpq.compare(gpq.data[leftInd], gpq.data[largest]) {
//				largest = leftInd
//			}
//
//			if rightInd < gpq.length && gpq.compare(gpq.data[rightInd], gpq.data[largest]) {
//				largest = rightInd
//			}
//
//			if largest == i {
//				break
//			}
//
//			gpq.data[i], gpq.data[largest] = gpq.data[largest], gpq.data[i]
//			i = largest
//		}
//	}
func (gpq *GenericPriorityQueue[T]) Peek() T {
	return gpq.data[0]
}

func Main1() {
	type Item struct {
		score1 int
		score2 int
	}

	gpq := NewGenericPriorityQueue[Item](func(t1, t2 Item) bool {
		return t1.score1 < t2.score1
	})
	gpq.Push(Item{1, 0})
	fmt.Println(gpq)
	gpq.Push(Item{-1, 0})
	fmt.Println(gpq)
	gpq.Push(Item{3, 0})
	fmt.Println(gpq)
	gpq.Push(Item{2, 0})
	fmt.Println(gpq)
	gpq.Push(Item{1, 0})
	fmt.Println(gpq)
	gpq.Push(Item{-2, 0})
	fmt.Println(gpq)
	fmt.Println(gpq.Pop())
	fmt.Println(gpq.Pop())
	fmt.Println(gpq.Pop())
}
