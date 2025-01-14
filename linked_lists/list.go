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

func (l *List[T]) print() {
	n := l.head
	for n.nextNode != nil {
		fmt.Println(n.value)
		n = n.nextNode
	}
	fmt.Println(n.value)
}
