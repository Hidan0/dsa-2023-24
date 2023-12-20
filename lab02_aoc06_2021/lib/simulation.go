package lib

import (
	"fmt"
)

type Simulation struct {
	f map[int]int
}

func addOr(m *map[int]int, at int, amount int, or int) {
	if (*m)[at] == 0 {
		(*m)[at] = or
	} else {
		(*m)[at] += amount
	}
}

func NewSimulation(initial []int) Simulation {
	var fishes map[int]int = make(map[int]int)

	for _, fish := range initial {
		addOr(&fishes, fish, 1, 1)
	}

	return Simulation{f: fishes}
}

const INITIAL_FISH = 6
const NEW_FISH = 8

func (s *Simulation) NextStep() {
	var newF map[int]int = make(map[int]int)

	for age, amount := range s.f {
		age--

		if age < 0 {
			addOr(&newF, INITIAL_FISH, amount, amount)
			addOr(&newF, NEW_FISH, amount, amount)
		} else {
			addOr(&newF, age, amount, amount)
		}
	}

	s.f = newF
}

func (s *Simulation) Steps(stp int) {
	for i := 0; i < stp; i++ {
		s.NextStep()
	}
}

func (s *Simulation) PrintState() string {
	var out string = ""

	for age, amount := range s.f {
		out += fmt.Sprintf("(%d * %d)", amount, age)
	}

	return out
}

func (s *Simulation) TotalFishes() int {
	tot := 0

	for _, amount := range s.f {
		tot += amount
	}

	return tot
}
