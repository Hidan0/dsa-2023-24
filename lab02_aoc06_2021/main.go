package main

import (
	"bufio"
	"fmt"
	"io"
	"lanternfish/lib"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	sim, err := parseInput(f)

	if err != nil {
		fmt.Println("Error parsing file:", err)
	}

	sim.Steps(256)
	fmt.Println("Total fishes after 256 days: ", sim.TotalFishes())
}

func parseInput(r io.Reader) (*lib.Simulation, error) {
	scanner := bufio.NewScanner(r)

	initial := make([]int, 0)

	if scanner.Scan() {
		src := strings.Split(scanner.Text(), ",")

		for _, raw := range src {
			fish, err := strconv.Atoi(raw)

			if err != nil {
				return nil, err
			}

			initial = append(initial, fish)
		}

	}

	sim := lib.NewSimulation(initial)
	return &sim, nil
}
