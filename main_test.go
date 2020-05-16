package main

import (
	"calc/point"
	"testing"
)

func TestDetermineMinimum(t *testing.T) {
	pA := point.Point{
		X: 1,
		Y: 1,
	}
	pB := point.Point{
		X: 1,
		Y: 3,
	}
	ans := point.DetermineMinimum(pA, pB)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}

func TestDetermineMinimumOfArray(t *testing.T) {
	pA := point.Point{
		X: 1,
		Y: 1,
	}
	pB := point.Point{
		X: 1,
		Y: 3,
	}
	pC := point.Point{
		X: 4,
		Y: 4,
	}
	pD := point.Point{
		X: 6,
		Y: 1,
	}

	ans := point.DetermineMinimumOfArray(pA, pB, pC, pD)
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}
