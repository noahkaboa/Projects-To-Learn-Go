package main

import (
	"fmt"
)

func main() {
	g := Graph[string]{}

	printGraph(g)

	nodeA := g.addNode("A!")
	nodeB := g.addNode("B!!")

	printGraph(g)

	g.addEdgeWithNodes(nodeA, nodeB)

	printGraph(g)
}

func printGraph[T any](g Graph[T]) {
	fmt.Println("Nodes:")
	for i := 0; i < len(g.nodes); i++ {
		fmt.Println(&g.nodes[i])
	}
	fmt.Println("Edges:")
	for i := 0; i < len(g.edges); i++ {
		fmt.Println(&g.edges[i])
	}

}
