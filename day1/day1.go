package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x, y      int
	direction string
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (loc *point) Move(command string) {
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

func Step1(input string) int {
	loc := point{0, 0, "N"}
	for _, cmd := range strings.Split(input, ", ") {
		loc.Move(cmd)
	}
	return Abs(loc.x) + Abs(loc.y)
}

func Step2(input string) int {

	return 0
}

func main() {
	content, err := ioutil.ReadFile("data/step1.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	fmt.Printf("Step1: %d\n", Step1(text))
}
