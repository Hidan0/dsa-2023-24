package lib

import (
	"strings"
	"testing"
)

const INPUT = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestOneFold(t *testing.T) {
	paper := InitPaper(strings.NewReader(INPUT))
	paper.Fold()

	want := 17
	res := paper.VisiblePoints()

	if res != want {
		t.Errorf("VisiblePoints() = %d, want %d", res, want)
	}
}

func TestFoldAll(t *testing.T) {
	paper := InitPaper(strings.NewReader(INPUT))
	paper.FoldAll()

	want := 16
	res := paper.VisiblePoints()

	if res != want {
		t.Errorf("VisiblePoints() = %d, want %d", res, want)
	}
}
