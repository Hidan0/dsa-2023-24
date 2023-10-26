package increasingsequences

import (
	"testing"
)

func equals(a, b *[][]int) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i, x := range *a {
		for j, y := range x {
			if y != (*b)[i][j] {
				return false
			}
		}
	}
	return true
}

func testNotEmpty(t *testing.T, values []int, want [][]int) {
	res, err := IncreasingSequences(&values)

	if err != nil {
		t.Fatalf(`Can not compute input due to: %v`, err)
	}

	if !equals(&want, &res) {
		t.Fatalf(`IncreasingSequences(%v) = %v, want %v`, values, res, want)
	}

}

func TestEmpty(t *testing.T) {
	values := []int{}

	_, err := IncreasingSequences(&values)
	if err == nil {
		t.Fatal("Should fail on empty input")
	}
}

func Input1(t *testing.T) {
	values := []int{9, 1, 3, 5, 2, 0, 8, 6}
	want := [][]int{{1, 3, 5}, {0, 8}}

	testNotEmpty(t, values, want)
}

func Test2(t *testing.T) {
	values := []int{9, 1, 3, 5, 2, 0, 8, 9, 6, 7, 7, 0}
	want := [][]int{{1, 3, 5}, {0, 8, 9}, {6, 7}}

	testNotEmpty(t, values, want)
}
