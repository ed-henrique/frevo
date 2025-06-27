package queues

import (
	"container/heap"
	"time"

	"github.com/ed-henrique/frevo/internal/assert"
	"github.com/ed-henrique/frevo/pkg/event"
)

type EventQueue struct {
	pq PriorityQueue
}

func NewEventQueue() *EventQueue {
	pq := make(PriorityQueue, 0)
	return &EventQueue{pq}
}

func (q *EventQueue) Len() int {
	return q.pq.Len()
}

func (q *EventQueue) Push(e event.Event, timestamp time.Duration) {
	item := &Item{value: e, timestamp: timestamp}
	heap.Push(&q.pq, item)
}

func (q *EventQueue) Pop() (event.Event, time.Duration) {
	assert.AssertTrue(q.pq.Len() > 0, errEmptyPriorityQueue)
	item := heap.Pop(&q.pq).(*Item)
	return item.value.(event.Event), item.timestamp
}
