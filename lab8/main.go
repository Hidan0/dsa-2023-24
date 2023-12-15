package main

import (
	"fmt"
	"lab8/lib"
)

func main() {
	lst := new(lib.LinkedList)
	lst.PushNode(1)
	fmt.Println(lst.Print())
	lst.PushNode(2)
	lst.PushNode(3)
	lst.PushNode(4)
	lst.PushNode(50)
	lst.PushNode(25)
	fmt.Println(lst.Print())
}
