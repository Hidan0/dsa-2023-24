package main

import (
	"bufio"
	"errors"
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

func (hm *HeightMap) isLowPoint(x, y int) bool {
	target := hm.GetAt(x, y)
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

func xyFromIdx(idx, width int) (int, int) {
	return idx % width, idx / width
}

func (hm *HeightMap) LowPoints() []int {
	var out []int

	hm.lowPoints(0, &out)

	return out
}

func (hm *HeightMap) lowPoints(idx int, lowPts *[]int) {
	if idx >= len(hm.values) {
		return
	}
	x, y := xyFromIdx(idx, hm.Width)
	if hm.isLowPoint(x, y) {
		(*lowPts) = append(*lowPts, idx)
	}
	hm.lowPoints(idx+1, lowPts)
}

func (hm *HeightMap) RiskLevel() uint {
	lowPoints := hm.LowPoints()

	return hm.riskLevel(0, &lowPoints)
}

func (hm *HeightMap) riskLevel(idx int, lowPoints *[]int) uint {
	if idx >= len(*lowPoints) {
		return 0
	}

	return 1 + uint(hm.values[(*lowPoints)[idx]]) + hm.riskLevel(idx+1, lowPoints)
}

func (hm *HeightMap) ThreeLargestBasins() uint {
	return 0
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
