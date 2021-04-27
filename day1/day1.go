package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type walker struct {
	loc       point
	direction string
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Step(dir string) point {
	s := map[string]point{
		"N": {1, 0},
		"E": {0, 1},
		"S": {0, -1},
		"W": {-1, 0},
	}
	return s[dir]
}

func Turn(orig string, side string) string {
	left := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	right := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}
	if side == "L" {
		return left[orig]
	} else {
		return right[orig]
	}
}

func (p *point) Dist() int {
	return Abs(p.x) + Abs(p.y)
}

func (w *walker) Move(command string) {
	var dir = command[0:1]
	distance, _ := strconv.Atoi(command[1:])
	w.direction = Turn(w.direction, dir)
	step := Step(w.direction)
	// fmt.Printf("w: %v  cmd: <%s %d>  step: %v\n", w, dir, distance, step)
	for i := 0; i < distance; i++ {
		w.loc.x += step.x
		w.loc.y += step.y
	}
}

func (w *walker) desc() {
	fmt.Printf("%v\n", w)
}

func Part1(input string) int {
	w := walker{point{0, 0}, "N"}
	for _, cmd := range strings.Split(strings.TrimRight(input, "\r\n"), ", ") {
		w.Move(cmd)
	}
	return w.loc.Dist()
}

func Part2(input string) int {
	previous := make(map[point]bool)
	w := walker{point{0, 0}, "N"}
	// fmt.Printf("%v dist %d\n", w, w.loc.Dist())
	previous[w.loc] = true
	for _, cmd := range strings.Split(strings.TrimRight(input, "\r\n"), ", ") {
		w.Move(cmd)
		// fmt.Printf("%v dist %d after %s\n", w, w.loc.Dist(), cmd)

		if previous[w.loc] {
			return w.loc.Dist()
		}
		previous[w.loc] = true
	}
	return -1
}

func main() {
	content, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	fmt.Printf("Part1: %d\n", Part1(text))
	fmt.Printf("Part2: %d\n", Part2(text))
}
