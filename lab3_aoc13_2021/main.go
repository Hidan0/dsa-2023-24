package main

import (
	"fmt"
	"os"
	"transparent_origami/lib"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	paper := lib.InitPaper(f)

	paper.Fold()

	fmt.Println(paper.VisiblePoints())
}
