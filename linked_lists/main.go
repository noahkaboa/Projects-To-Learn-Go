package main

import (
	"fmt"
)

func main() {
	list := new(List[string])
	list.add("Hello!")
	list.add("World!")

	fmt.Println("Two elements added")
	// list.print()

	fmt.Println(list.get(-1))

}
