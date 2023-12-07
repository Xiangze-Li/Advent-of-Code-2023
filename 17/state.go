package puzzle17

import (
	c "advent2023/util/constants"
	"container/heap"
)

type state struct {
	i, j  int
	going []c.Direction
	cost  uint64
}

type stateHeap []state

func (a stateHeap) Len() int           { return len(a) }
func (a stateHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a stateHeap) Less(i, j int) bool { return a[i].cost < a[j].cost }
func (a *stateHeap) Push(elem any)     { *a = append(*a, elem.(state)) }
func (a *stateHeap) Pop() any {
	length := len(*a)
	elem := (*a)[length-1]
	*a = (*a)[:length-1]
	return elem
}

var _ heap.Interface = (*stateHeap)(nil)
