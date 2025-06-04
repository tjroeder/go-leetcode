package slidingwindowmaximum

import (
	"container/list"
	"strconv"
)

type Deque struct {
	items *list.List
}

// NewDeque is a constructor that will declare and return the Deque type object
func NewDeque() *Deque {
	return &Deque{list.New()}
}

// PushFront will push an element at the front of the dequeue
func (d *Deque) PushFront(val int) {
	d.items.PushFront(val)
}

// PushBack will push an element at the back of the dequeue
func (d *Deque) PushBack(val int) {
	d.items.PushBack(val)
}

// PopFront will pop an element from the front of the dequeue
func (d *Deque) PopFront() int {
	return d.items.Remove(d.items.Front()).(int)
}

// PopBack will pop an element from the back of the dequeue
func (d *Deque) PopBack() int {
	return d.items.Remove(d.items.Back()).(int)
}

// Front will return the element from the front of the dequeue
func (d *Deque) Front() int {
	return d.items.Front().Value.(int)
}

// Back will return the element from the back of the dequeue
func (d *Deque) Back() int {
	return d.items.Back().Value.(int)
}

// Empty will check if the dequeue is empty or not
func (d *Deque) Empty() bool {
	return d.items.Len() == 0
}

// Len will return the length of the dequeue
func (d *Deque) Len() int {
	return d.items.Len()
}

func (d *Deque) Print() string {
	temp := d.items.Front()
	s := "["
	for temp != nil {
		temp2, _ := temp.Value.(int)
		s += strconv.Itoa(temp2)
		temp = temp.Next()
		if temp != nil {
			s += " , "
		}
	}
	s += "]"
	return s
}
