package main

import "fmt"

type Node[T any] struct {
	value    T
	nextNode *Node[T]
}

type List[T any] struct {
	head *Node[T]
}

func (l *List[T]) add(element T) {
	if l.head == nil {
		l.head = &Node[T]{value: element, nextNode: nil}
		return
	}
	n := l.head
	for n.nextNode != nil {
		n = n.nextNode
	}
	n.nextNode = &Node[T]{value: element, nextNode: nil}
	return
}

func (l *List[T]) get(index int) T {

	if index < 0 {
		return nothing[T]()
	}

	n := l.head
	for i := 0; i < index; i += 1 {
		if n.nextNode == nil {
			return nothing[T]()
		}
		n = n.nextNode
	}
	return n.value
}

func (l *List[T]) print() {
	n := l.head
	for n.nextNode != nil {
		fmt.Println(n.value)
		n = n.nextNode
	}
	fmt.Println(n.value)
}

func nothing[T any]() T {
	var zero T
	return zero
}
