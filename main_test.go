package main

import (
  "testing"
  "math"
)

func TestDetermineMinimumAB(t *testing.T) {
	pA := Point{
		X: 1,
		Y: 1,
	}
	pB := Point{
		X: 1,
		Y: 3,
	}
	ans := DetermineMinimum(pA, pB)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}

func TestDetermineMinimumAC(t *testing.T) {
	pA := Point{
		X: 1,
		Y: 1,
	}
	pC := Point{
		X: 4,
		Y: 4,
  }
  ans := DetermineMinimum(pA, pC)
	if math.Round(ans) != 4.0 {
		t.Errorf("DetermineMinimum()= %f; want 4.242641", ans)
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) > float64EqualityThreshold
}

func TestDetermineMinimumOfArray(t *testing.T) {
	pA := Point{
		X: 1,
		Y: 1,
	}
	pB := Point{
		X: 1,
		Y: 3,
	}
	pC := Point{
		X: 4,
		Y: 4,
	}
	pD := Point{
		X: 6,
		Y: 1,
	}

	ans := DetermineMinimumOfArray(pA, pB, pC, pD)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}
