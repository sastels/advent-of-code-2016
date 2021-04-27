package main

import (
	"testing"
)

func TestStep1(t *testing.T) {
	var loc = point{0, 0, "N"}
	loc.Step("L1")
	var expected = point{-1, 0, "W"}
	if loc != expected {
		t.Errorf("loc = %v; want %v", loc, expected)
	}
}

func TestStep2(t *testing.T) {
	var loc = point{10, 30, "E"}
	loc.Step("R10")
	var expected = point{10, 20, "S"}
	if loc != expected {
		t.Errorf("loc = %v; want %v", loc, expected)
	}
}
