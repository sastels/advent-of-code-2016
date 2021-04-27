package main

import (
	"testing"
)

func TestMove1(t *testing.T) {
	var loc = point{0, 0, "N"}
	loc.Move("L1")
	var expected = point{-1, 0, "W"}
	if loc != expected {
		t.Errorf("loc = %v; want %v", loc, expected)
	}
}

func TestMove2(t *testing.T) {
	var loc = point{10, 30, "E"}
	loc.Move("R10")
	var expected = point{10, 20, "S"}
	if loc != expected {
		t.Errorf("loc = %v; want %v", loc, expected)
	}
}

func TestStep1(t *testing.T) {
	actual := Step1("R5, L5, R5, R3")
	expected := 12

	if actual != expected {
		t.Errorf("actual %d expected %d", actual, expected)
	}

}
