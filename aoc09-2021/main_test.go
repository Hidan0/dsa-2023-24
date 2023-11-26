package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	want := `2199943210
3987894921
9856789892
8767896789
9899965678`

	hm, err := parseInput(bufio.NewReader(strings.NewReader(want)))
	if err != nil {
		t.Error(err)
	}

	if hm.Width != 10 {
		t.Errorf("Width should be 10, got %d", hm.Width)
	}

	if hm.Height != 5 {
		t.Errorf("Height should be 5, got %d", hm.Height)
	}

	if !slices.Equal(hm.values, []uint8{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}) {
		t.Errorf("Values parsed incorrectly, got %v", hm.values)
	}
}

func TestRiskLevel(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	hm, err := parseInput(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Error(err)
	}

	var want uint = 15
	res := hm.RiskLevel()

	if res != want {
		t.Errorf("RiskLevel() = %d, expected %d", res, want)
	}
}

func TestLowPointsOnInputFile(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	hm, err := parseInput(bufio.NewReader(f))
	if err != nil {
		t.Error(err)
	}

	var want uint = 570
	res := hm.RiskLevel()

	if res != want {
		t.Errorf("RiskLevel() = %d, expected %d", res, want)
	}
}

func TestBasin(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	hm, err := parseInput(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Error(err)
	}
	var want uint = 1134
	res := hm.ThreeLargestBasins()

	if res != want {
		t.Errorf("Basins() = %d, expected %d", res, want)
	}
}

func TestBasinOnInputFile(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	hm, err := parseInput(bufio.NewReader(f))
	if err != nil {
		t.Error(err)
	}

	var want uint = 899392
	res := hm.ThreeLargestBasins()

	if res != want {
		t.Errorf("Basins() = %d, expected %d", res, want)
	}
}
