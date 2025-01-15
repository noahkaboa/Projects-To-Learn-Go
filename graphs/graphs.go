package main

type Node[T any] struct {
	value T
}

type Edge[T any] struct {
	NodeA *Node[T]
	NodeB *Node[T]
}

type Graph[T any] struct {
	nodes []Node[T]
	edges []Edge[T]
}

func (g *Graph[T]) addNode(value T) Node[T] {
	newNode := Node[T]{value: value}
	g.nodes = append(g.nodes, newNode)
	return newNode
}

func (g *Graph[T]) addEdgeWithNodes(nodeA Node[T], nodeB Node[T]) {
	var newEdge Edge[T]
	newEdge.NodeA = &nodeA
	newEdge.NodeB = &nodeB
	g.edges = append(g.edges, newEdge)
}
