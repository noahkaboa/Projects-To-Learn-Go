package main

import (
	"fmt"
)

func main() {
	g := Graph[string]{}

	printGraph(g)

	nodeAPtr := g.addNode("A!")
	nodeBPtr := g.addNode("B!!")

	printGraph(g)

	g.addEdgeWithNodes(nodeAPtr, nodeBPtr)

	printGraph(g)

	g.nodes[0].value = "Not A!"

	printGraph(g)

	nodeCPtr := g.addNode("its c")
	nodeDPtr := g.addNode("DDD")

	printGraph(g)

	g.addEdgeWithNodes(nodeCPtr, nodeDPtr)

	printGraph(g)

	g.addEdgeWithNodes(nodeAPtr, nodeCPtr)
	printGraph(g)
}

func printGraph[T any](g Graph[T]) {
	fmt.Println("Printing graph!")
	fmt.Println("Nodes:")
	for i := 0; i < len(g.nodes); i++ {
		fmt.Println(g.nodes[i])
	}
	fmt.Println("Edges:")
	for i := 0; i < len(g.edges); i++ {
		fmt.Print("NodeA:\t")
		fmt.Print(*g.edges[i].NodeAPtr)
		fmt.Print("\t\tNodeB:\t")
		fmt.Println(*g.edges[i].NodeBPtr)
	}

}
