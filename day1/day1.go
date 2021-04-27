package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x, y      int
	direction string
}

func (loc *point) Step(command string) {
	var dir = command[0:1]
	if dir == "R" {
		switch loc.direction {
		case "N":
			loc.direction = "E"
		case "E":
			loc.direction = "S"
		case "S":
			loc.direction = "W"
		case "W":
			loc.direction = "N"
		}
	} else if dir == "L" {
		switch loc.direction {
		case "N":
			loc.direction = "W"
		case "W":
			loc.direction = "S"
		case "S":
			loc.direction = "E"
		case "E":
			loc.direction = "N"
		}
	} else {
		fmt.Printf("bad direction %s!!\n", dir)
	}

	distance, error := strconv.Atoi(command[1:])
	if error == nil {
		switch loc.direction {
		case "N":
			loc.y += distance
		case "E":
			loc.x += distance
		case "S":
			loc.y -= distance
		case "W":
			loc.x -= distance
		}
	}
}

func main() {
	var location = point{0, 0, "N"}
	fmt.Printf("%v\n", location)

	location.Step("L4")

	fmt.Printf("%v\n", location)
}
