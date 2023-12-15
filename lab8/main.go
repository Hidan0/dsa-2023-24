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
}
