package lanternfish

import (
	"strconv"
	"strings"
)

type Simulation struct {
	f []int
}

func NewSimulation(initial []int) Simulation {
	return Simulation{f: initial}
}

const INITIAL_FISH = 6
const NEW_FISH = 8

func (s *Simulation) updateFishAt(i int) {
	fish := &(s.f[i])

	(*fish)--

	if *fish < 0 {
		*fish = INITIAL_FISH
		s.f = append(s.f, NEW_FISH)
	}

}

func (s *Simulation) NextStep() {
	for i := range s.f {
		s.updateFishAt(i)
	}
}

func (s *Simulation) Steps(stp int) {
	for i := 0; i < stp; i++ {
		s.NextStep()
	}
}

func (s *Simulation) PrintState() string {
	var str []string

	for _, fish := range s.f {
		str = append(str, strconv.Itoa(fish))
	}

	return strings.Join(str, ",")
}

func (s *Simulation) TotalFishes() int {
	return len(s.f)
}
