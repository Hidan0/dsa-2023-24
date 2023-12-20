package main

// USE RECURSION

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnergyLevel struct {
	Width  int
	values []uint8
}

func (e *EnergyLevel) xyFromIdx(idx int) (int, int) {
	return idx % e.Width, idx / e.Width
}

func appendIfNotPresent(slice *[]int, element int) {
	for _, e := range *slice {
		if e == element {
			return
		}
	}
	*slice = append(*slice, element)
}

func parseInput(reader *bufio.Reader) (*EnergyLevel, error) {
	Width := -1
	var values []uint8

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			break
		}

		if Width < 0 {
			Width = len(line)
		}

		for _, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			values = append(values, uint8(value))
		}
	}

	if Width < 0 {
		return nil, errors.New("Invalid input")
	}

	hm := EnergyLevel{Width, values}
	return &hm, nil
}

func (e *EnergyLevel) toString() string {
	var sb strings.Builder
	for i, value := range e.values {
		sb.WriteString(strconv.Itoa(int(value)))
		if (i+1)%e.Width == 0 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (e *EnergyLevel) NextStep() uint {
	var totalFlashed uint = 0

	var flashed []bool = make([]bool, len(e.values))
	var toFlash []int

	e.incrementByOne(0, &toFlash)
	e.flash(0, &toFlash, &flashed, &totalFlashed)

	return totalFlashed
}

func (e *EnergyLevel) Advance(steps uint) uint {
	var out uint
	for i := uint(0); i < steps; i++ {
		out += e.NextStep()
	}
	return out
}

// Increment recursively by one the values of the energy level.
// If the value is greater than 9, add the index to the list of indexes to flash.
func (e *EnergyLevel) incrementByOne(idx int, toFlash *[]int) {
	if idx >= len(e.values) {
		return
	}

	e.values[idx] += 1

	if e.values[idx] > 9 {
		(*toFlash) = append((*toFlash), idx)
	}

	e.incrementByOne(idx+1, toFlash)
}

func (e *EnergyLevel) flash(idx int, toFlash *[]int, flashed *[]bool, totalFlashed *uint) {
	if idx >= len(*toFlash) {
		return
	}

	targetIdx := (*toFlash)[idx]

	if (*flashed)[targetIdx] || e.values[targetIdx] <= 9 {
		return
	}

	e.values[targetIdx] = 0
	(*flashed)[targetIdx] = true
	*totalFlashed += 1

	x, y := e.xyFromIdx(targetIdx)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < e.Width && ny >= 0 && ny < e.Width {
				if !(*flashed)[ny*e.Width+nx] {
					e.values[ny*e.Width+nx] += 1

					if e.values[ny*e.Width+nx] > 9 {
						appendIfNotPresent(toFlash, ny*e.Width+nx)
					}
				}
			}
		}
	}

	e.flash(idx+1, toFlash, flashed, totalFlashed)
}

func (e *EnergyLevel) isSync() bool {
	for _, value := range e.values {
		if value > 0 {
			return false
		}
	}
	return true
}

func (e *EnergyLevel) Sync() int {
	const MAX_ITERATIONS = 1000

	i := 0
	for !e.isSync() && i < MAX_ITERATIONS {
		e.NextStep()
		i++
	}

	if i == MAX_ITERATIONS {
		return -1
	}
	return i
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	e, err := parseInput(bufio.NewReader(f))
	if err != nil {
		fmt.Println(err)
		return
	}
	// res := e.Advance(100)
	// fmt.Println(res)
	//
	fmt.Println(e.Sync())
}
