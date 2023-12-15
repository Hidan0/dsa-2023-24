package main

import (
	"fmt"
	"lab8/lib"
)

func main() {
	lst := new(lib.LinkedList)
	fmt.Println("Search 2 ", lst.Search(2))
	lst.PushNode(1)
	fmt.Println(lst.Print())
	fmt.Println("Search 2 ", lst.Search(2))
	lst.PushNode(2)
	lst.PushNode(3)
	lst.PushNode(4)
	lst.PushNode(50)
	lst.PushNode(25)
	fmt.Println(lst.Print())

	fmt.Println("Search 25 ", lst.Search(25))
	fmt.Println("Search 1 ", lst.Search(1))
	fmt.Println("Search 3 ", lst.Search(3))

	fmt.Println("Remove -1 ", lst.Remove(-1))
	fmt.Println("Remove 1 ", lst.Remove(1))
	fmt.Println("Remove 3 ", lst.Remove(3))
	fmt.Println("Remove 25 ", lst.Remove(25))
	fmt.Println("Remove 25 ", lst.Remove(25))
}
