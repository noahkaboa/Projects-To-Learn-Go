package main

type Node[T any] struct {
	value    T
	nextNode *Node[T]
}

type List[T any] struct {
	head *Node[T]
}

func (l List[T]) add(element T) {
	n := l.head
	for n != nil {
		n = n.nextNode
	}
	n = &Node[T]{value: element, nextNode: nil}
	return
}
