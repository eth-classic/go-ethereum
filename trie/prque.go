package trie

import (
	"container/heap"
)

// Priority queue data structure.
type Prque struct {
	cont *sstack
}

// Creates a new priority queue.
func NewPrque() *Prque {
	return &Prque{newSstack()}
}

// Pushes a value with a given priority into the queue, expanding if necessary.
func (p *Prque) Push(data interface{}, priority float32) {
	heap.Push(p.cont, &item{data, priority})
}

// Pops the value with the greates priority off the stack and returns it.
// Currently no shrinking is done.
func (p *Prque) Pop() (interface{}, float32) {
	item := heap.Pop(p.cont).(*item)
	return item.value, item.priority
}

// Pops only the item from the queue, dropping the associated priority value.
func (p *Prque) PopItem() interface{} {
	return heap.Pop(p.cont).(*item).value
}

// Checks whether the priority queue is empty.
func (p *Prque) Empty() bool {
	return p.cont.Len() == 0
}

// Returns the number of element in the priority queue.
func (p *Prque) Size() int {
	return p.cont.Len()
}

// Clears the contents of the priority queue.
func (p *Prque) Reset() {
	*p = *NewPrque()
}

