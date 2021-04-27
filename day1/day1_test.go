package main

import (
	"testing"
)

func TestMove1(t *testing.T) {
	w := walker{point{0, 0}, "N"}
	w.Move("L1")
	expected := walker{point{-1, 0}, "W"}
	if w != expected {
		t.Errorf("loc = %v; want %v", w, expected)
	}
}

func TestMove2(t *testing.T) {
	w := walker{point{10, 30}, "E"}
	w.Move("R10")
	expected := walker{point{10, 20}, "S"}
	if w != expected {
		t.Errorf("w = %v; w %v", w, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1("R5, L5, R5, R3")
	expected := 12
	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}
}
