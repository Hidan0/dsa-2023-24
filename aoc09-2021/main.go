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
