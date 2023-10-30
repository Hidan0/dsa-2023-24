package gatherpermutations

import "testing"

func testPermutations(t *testing.T, a []int, want int) {
	res := GatherPermutations(&a)

	if res != want {
		t.Errorf("GatherPermutations(%v) = %d, want %d", a, res, want)
	}
}

func TestGivenInput(t *testing.T) {
	a := []int{4, 5, 1, 3, 6, 2}
	testPermutations(t, a, 2)
}

func TestAllForward(t *testing.T) {
	a := []int{1, 2, 3, 4}
	testPermutations(t, a, 0)
}

func TestAllBack(t *testing.T) {
	a := []int{4, 3, 2, 1}
	testPermutations(t, a, 3)
}

func TestRandom(t *testing.T) {
	a := []int{3, 5, 1, 2, 4}
	testPermutations(t, a, 2)
}
