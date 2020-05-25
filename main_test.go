package main

import (
  "testing"
  "math"
)

var pA Point = Point{
  X: 1,
  Y: 1,
}
var pB Point = Point{
  X: 1,
  Y: 3,
}
var pC Point = Point{
  X: 4,
  Y: 4,
}
var pD Point = Point{
  X: 6,
  Y: 1,
}

func TestDetermineMinimumAB(t *testing.T) {
	ans := DetermineMinimum(pA, pB)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}

func TestDetermineMinimumAC(t *testing.T) {
  ans := DetermineMinimum(pA, pC)
	if math.Round(ans) != 14.0 {
		t.Errorf("DetermineMinimum()= %f; want 4.242641", ans)
	}
}

func TestDetermineMinimumOfArray(t *testing.T) {
  points := []Point{pA, pB, pC, pD}
	ans := DetermineMinimumOfArray(points)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}
