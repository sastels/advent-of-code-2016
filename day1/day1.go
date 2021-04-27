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

func (p *point) Dist() int {
	return Abs(p.x) + Abs(p.y)
}

func (w *walker) Move(command string) {
	var dir = command[0:1]
	if dir == "R" {
		switch w.direction {
		case "N":
			w.direction = "E"
		case "E":
			w.direction = "S"
		case "S":
			w.direction = "W"
		case "W":
			w.direction = "N"
		}
	} else if dir == "L" {
		switch w.direction {
		case "N":
			w.direction = "W"
		case "W":
			w.direction = "S"
		case "S":
			w.direction = "E"
		case "E":
			w.direction = "N"
		}
	} else {
		fmt.Printf("bad direction %s!!\n", dir)
	}

	distance, error := strconv.Atoi(command[1:])
	if error == nil {
		switch w.direction {
		case "N":
			w.loc.y += distance
		case "E":
			w.loc.x += distance
		case "S":
			w.loc.y -= distance
		case "W":
			w.loc.x -= distance
		}
	}
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
	previous[w.loc] = true
	for _, cmd := range strings.Split(strings.TrimRight(input, "\r\n"), ", ") {
		w.Move(cmd)
		if previous[w.loc] {
			return w.loc.Dist()
		}
		previous[w.loc] = true
	}
	return -1
}

func main() {
	content, err := ioutil.ReadFile("data/step1.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	fmt.Printf("Part1: %d\n", Part1(text))
	fmt.Printf("Part2: %d\n", Part2(text))
}
