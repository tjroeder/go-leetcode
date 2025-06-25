package ipo

import "container/heap"

// use min heap for finding the minimum capital, and max heap for the maximum profit
func findMaximizedCapital(k int, w int, profits []int, capitals []int) int {
	// set currCap to our starting capital, and initialize the min and max heaps
	currCap := w
	capitalsMinHeap := newMinHeap()
	profitsMaxHeap := newMaxHeap()

	// push the capitals into the min heap, use a set to define the capitals value with index
	for i, capital := range capitals {
		heap.Push(capitalsMinHeap, set{index: i, value: capital})
	}

	// outer loop to select k projects
	for i := 0; i < k; i++ {
		// loop through our capitals min heap
		// if our min heap is not empty, and the root of the capitals min heap value
		// is less than or equal to the current capital value
		for len(*capitalsMinHeap) != 0 && capitalsMinHeap.Top().(set).value <= currCap {
			// we pop the capital set from the heap (our project)
			project := heap.Pop(capitalsMinHeap).(set)
			u := project.index
			// using the index of the project, we push the value into the max heap
			heap.Push(profitsMaxHeap, set{value: profits[u]})
		}

		// if we never pushed anything in, in the first place we know that we don't
		// have enough starting capital
		if len(*profitsMaxHeap) == 0 {
			break
		}

		// when we pushed project profit to the heap, we need to adjust our current
		// capital value from taking profits
		j := heap.Pop(profitsMaxHeap).(set).value
		currCap += j
	}
	return currCap
}

type set struct {
	index int
	value int
}

type minHeap []set

func newMinHeap() *minHeap {
	min := &minHeap{}
	heap.Init(min)
	return min
}

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(set))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h minHeap) Top() interface{} {
	return h[0]
}

type maxHeap []set

func newMaxHeap() *maxHeap {
	max := &maxHeap{}
	heap.Init(max)
	return max
}

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(set))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h maxHeap) Top() interface{} {
	return h[0]
}
