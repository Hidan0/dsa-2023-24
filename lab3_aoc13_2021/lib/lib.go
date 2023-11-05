package lib

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type command struct {
	op    byte
	coord int
}

type TPaper struct {
	points []point
	cmd    []command
}

func InitPaper(r io.Reader) *TPaper {
	scanner := bufio.NewScanner(r)

	p := []point{}
	cmd := []command{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "fold along") {
			op := strings.Split(strings.Split(line, " ")[2], "=")

			coord, err := strconv.Atoi(op[1])
			if err != nil {
				fmt.Println("Error converting coordinate to int")
				return nil
			}

			if op[0] == "x" || op[0] == "y" {
				cmd = append(cmd, command{op[0][0], coord}) // this is a bit ugly
			} else {
				fmt.Println("Error parsing fold command")
				return nil
			}

		} else {
			coords := strings.Split(line, ",")

			x, err := strconv.Atoi(coords[0])
			if err != nil {
				fmt.Printf("Error converting x: %s coordinate to int\n", coords[0])
				return nil
			}

			y, err := strconv.Atoi(coords[1])
			if err != nil {
				fmt.Printf("Error converting y: %s coordinate to int\n", coords[1])
				return nil
			}
			p = append(p, point{x, y})
		}
	}

	return &TPaper{p, cmd}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func appendUnique(p *[]point, pt point) {
	for _, p := range *p {
		if p.x == pt.x && p.y == pt.y {
			return
		}
	}
	*p = append(*p, pt)
}

func (p *TPaper) Fold() {
	if len(p.cmd) <= 0 {
		return
	}

	cmd := p.cmd[0]
	p.cmd = p.cmd[1:]

	newPts := []point{}

	if cmd.op == 'x' {
		for _, pt := range p.points {
			newX := cmd.coord - abs(pt.x-cmd.coord)
			appendUnique(&newPts, point{newX, pt.y})
		}
	} else if cmd.op == 'y' {
		for _, pt := range p.points {
			newY := cmd.coord - abs(pt.y-cmd.coord)
			appendUnique(&newPts, point{pt.x, newY})
		}
	}

	p.points = newPts
}

func (p *TPaper) FoldAll() {
	for len(p.cmd) > 0 {
		p.Fold()
	}
}

func (p *TPaper) VisiblePoints() int {
	return len(p.points)
}

func (p *TPaper) FindPt(x, y int) string {
	for _, pt := range p.points {
		if pt.x == x && pt.y == y {
			return "#"
		}
	}
	return "."
}

func (p *TPaper) Print() string {
	out := ""
	maxX, maxY := 0, 0

	for _, pt := range p.points {
		if pt.x > maxX {
			maxX = pt.x
		}
		if pt.y > maxY {
			maxY = pt.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			out += p.FindPt(x, y)
		}
		out += "\n"
	}

	return out
}
