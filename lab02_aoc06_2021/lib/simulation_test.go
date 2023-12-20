package lib

import (
	"strings"
	"testing"
)

func TestInitialPrintState(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := []string{"(1 * 1)", "(1 * 2)", "(2 * 3)", "(1 * 4)"}

	res := s.PrintState()
	for _, sub := range want {
		if !strings.Contains(res, sub) {
			t.Errorf("PrintState() = %v, does not contains %v", res, sub)
		}
	}
}

func TestOneStep(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := []string{"(1 * 0)", "(1 * 1)", "(2 * 2)", "(1 * 3)"}

	s.NextStep()
	res := s.PrintState()
	for _, sub := range want {
		if !strings.Contains(res, sub) {
			t.Errorf("PrintState() = %v, does not contains %v", res, sub)
		}
	}
}

func TestThreeSteps(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := []string{"(2 * 0)", "(1 * 1)", "(1 * 5)", "(1 * 6)", "(1 * 7)", "(1 * 8)"}

	s.Steps(3)

	res := s.PrintState()
	for _, sub := range want {
		if !strings.Contains(res, sub) {
			t.Errorf("PrintState() = %v, does not contains %v", res, sub)
		}
	}
}

func Test18Days(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := 26

	s.Steps(18)

	if s.TotalFishes() != want {
		t.Errorf("Steps(18) = %v, want %v", s.TotalFishes(), want)
	}
}

func Test80Days(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := 5934
	s.Steps(80)
	if s.TotalFishes() != want {
		t.Errorf("Steps(80) = %v, want %v", s.TotalFishes(), want)
	}
}
