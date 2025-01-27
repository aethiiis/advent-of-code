package utils

import "container/heap"

type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T any] []*Item[T]

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Priority > (*pq)[j].Priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}

func NewPriorityQueue[T any]() PriorityQueue[T] {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)
	return pq
}

func NewPriorityQueueFromList[T any](list []T) PriorityQueue[T] {
	pq := make(PriorityQueue[T], len(list))
	for i, value := range list {
		pq[i] = &Item[T]{
			Value:    value,
			Priority: i,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	return pq
}

func (pq *PriorityQueue[T]) Put(value T, priority int) {
	item := &Item[T]{Value: value, Priority: priority}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) Get() (T, int) {
	item := heap.Pop(pq).(*Item[T])
	return item.Value, item.Priority
}

func (pq *PriorityQueue[T]) Empty() bool {
	return len(*pq) == 0
}
