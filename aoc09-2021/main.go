package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HeightMap struct {
	Width  int
	Height int
	values []uint8
}

func (hm *HeightMap) Print() {
	for y := 0; y < hm.Height; y++ {
		for x := 0; x < hm.Width; x++ {
			print(hm.values[y*hm.Width+x])
		}
		println()
	}
}

func (hm *HeightMap) GetAt(x, y int) uint8 {
	return hm.values[y*hm.Width+x]
}

func (hm *HeightMap) isLowPoint(target uint8, x, y int) bool {
	if x > 0 && hm.GetAt(x-1, y) <= target {
		return false
	}
	if x < hm.Width-1 && hm.GetAt(x+1, y) <= target {
		return false
	}
	if y > 0 && hm.GetAt(x, y-1) <= target {
		return false
	}
	if y < hm.Height-1 && hm.GetAt(x, y+1) <= target {
		return false
	}

	return true
}

func (hm *HeightMap) LowPoints() uint {
	var out uint = 0
	hm.lowPoints(0, &out)
	return out
}

func (hm *HeightMap) lowPoints(i int, out *uint) {
	if i >= len(hm.values) {
		return
	}

	x := i % hm.Width
	y := i / hm.Width

	target := hm.GetAt(x, y)

	if hm.isLowPoint(target, x, y) {
		fmt.Printf("Found low point at (%d, %d) with value %d\n", x, y, target)
		*out += 1 + uint(target)
	}

	hm.lowPoints(i+1, out)
}

func parseInput(reader *bufio.Reader) (*HeightMap, error) {
	Height := 0
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
		Height += 1
	}

	if Width < 0 || Height == 0 {
		return nil, errors.New("Invalid input")
	}

	hm := HeightMap{Width, Height, values}
	return &hm, nil
}
