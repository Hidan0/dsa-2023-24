package lanternfish

import "testing"

func TestOneStep(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := "2,3,2,0,1"

	s.NextStep()

	if s.PrintState() != want {
		t.Errorf("NextStep() = %v, want %v", s.PrintState(), want)
	}
}

func TestThreeSteps(t *testing.T) {
	s := NewSimulation([]int{3, 4, 3, 1, 2})
	want := "0,1,0,5,6,7,8"

	s.NextStep()
	s.NextStep()
	s.NextStep()

	if s.PrintState() != want {
		t.Errorf("NextStep() = %v, want %v", s.PrintState(), want)
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
