package queues

import (
	"time"
)

const (
	errEmptyPriorityQueue = "the priority queue is empty."
)

// An Item is something we manage in a priority queue.
type Item struct {
	value any

	// At which moment the events will be triggered.
	timestamp time.Duration

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Len returns the length of the PriorityQueue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less defines how Items should be compared for the PriorityQueue.
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, timestamp so we use lesser than here.
	return pq[i].timestamp < pq[j].timestamp
}

// Less defines how Items should be swapped for the PriorityQueue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds a new Item to the PriorityQueue.
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop gets the Item with the lowest timestamp from the PriorityQueue.
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
