package main

import (
	"bufio"
	"fmt"
	"lab8/lib"
	"os"
	"strconv"
	"strings"
)

func main() {
	lst := new(lib.LinkedList)
	f, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, " ")
		if len(strs) == 2 {
			val, _ := strconv.Atoi(strs[1])
			switch strs[0] {
			case "+":
				insertIfNotExists(lst, val)
			case "-":
				removeIfExists(lst, val)
			case "?":
				if lst.Search(val) {
					fmt.Println("Found ", val)
				} else {
					fmt.Println("Not found ", val)
				}
			default:
				fmt.Println("Invalid command", strs[0])
			}
		} else {
			switch strs[0] {
			case "c":
				fmt.Println("Size:", lst.Size())
			case "p":
				fmt.Println(lst.Print())
			case "o":
				lst.PrintReversed()
			case "d":
				lst.Drop()
				fmt.Println("Dropped")
			case "f":
				return
			default:
				fmt.Println("Invalid command", strs[0])
			}
		}
	}
}

func insertIfNotExists(lst *lib.LinkedList, value int) {
	if !lst.Search(value) {
		lst.PushNode(value)
		fmt.Println("*** Inserted ", value)
	}
}

func removeIfExists(lst *lib.LinkedList, value int) {
	if lst.Remove(value) {
		fmt.Println("*** Removed ", value)
	}
}
