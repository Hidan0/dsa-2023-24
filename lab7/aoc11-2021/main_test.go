package main

import (
	"bufio"
	"strings"
	"testing"
)

const INPUT1 = `11111
19991
19191
19991
11111`

const INPUT2 = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestInput1OneStep(t *testing.T) {
	want := `34543
40004
50005
40004
34543
`

	e, err := parseInput(bufio.NewReader(strings.NewReader(INPUT1)))
	if err != nil {
		t.Errorf("parseInput() error = %v", err)
	}

	e.NextStep()
	res := e.toString()

	if res != want {
		t.Errorf("e.toString() = \n%v but want: \n%v", res, want)
	}
}

func TestInput1TwoSteps(t *testing.T) {
	want := `45654
51115
61116
51115
45654
`

	e, err := parseInput(bufio.NewReader(strings.NewReader(INPUT1)))
	if err != nil {
		t.Errorf("parseInput() error = %v", err)
	}

	e.NextStep()
	e.NextStep()
	res := e.toString()

	if res != want {
		t.Errorf("e.toString() = \n%v but want: \n%v", res, want)
	}
}

func TestInput2NumOfStepsAfter10(t *testing.T) {
	var want uint = 204

	e, err := parseInput(bufio.NewReader(strings.NewReader(INPUT2)))
	if err != nil {
		t.Errorf("parseInput() error = %v", err)
	}

	res := e.Advance(10)

	if res != want {
		t.Errorf("The total of flashes is %d, should be %d", res, want)
	}
}

func TestInput2NumOfStepsAfter100(t *testing.T) {
	var want uint = 1656

	e, err := parseInput(bufio.NewReader(strings.NewReader(INPUT2)))
	if err != nil {
		t.Errorf("parseInput() error = %v", err)
	}

	res := e.Advance(100)

	if res != want {
		t.Errorf("The total of flashes is %d, should be %d", res, want)
	}
}
