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

func (p *TPaper) Fold() {
	fmt.Println("FOLD")
}
