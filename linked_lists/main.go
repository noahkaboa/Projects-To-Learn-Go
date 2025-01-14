package main

import (
	"fmt"
)

func main() {
	list := new(List[string])
	list.add("Hello!")
	list.add("World!")

	fmt.Println("Two elements added")

	fmt.Println(list.get(-1))

}
