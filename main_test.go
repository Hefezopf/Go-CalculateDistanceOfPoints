package main

import (
	"math"
	"testing"
)

var pA Point = Point{X: 1, Y: 1}
var pB Point = Point{X: 1, Y: 3}
var pC Point = Point{X: 4, Y: 4}
var pD Point = Point{X: 6, Y: 1}

func TestParseAllCoords(t *testing.T) {
	var coords = "(4,3) (5,2) (1,1) (1,2)";

	ans := ParseAllCoords(coords)

	if len(ans) != 4 {
		t.Errorf("ParseToCoord()= %d; want 4", len(ans))
	}
    if ans[0].X != 4 {
		t.Errorf("ParseToCoord()= %d; want 4", ans[0].X)
	}    
    if ans[0].Y != 3 {
		t.Errorf("ParseToCoord()= %d; want 3", ans[0].Y)
	}    
    if ans[1].X != 5 {
		t.Errorf("ParseToCoord()= %d; want 5", ans[1].X)
	}    
    if ans[1].Y != 2 {
		t.Errorf("ParseToCoord()= %d; want 2", ans[1].Y)
	}    
    if ans[2].X != 1 {
		t.Errorf("ParseToCoord()= %d; want 1", ans[2].X)
	}    
    if ans[2].Y != 1 {
		t.Errorf("ParseToCoord()= %d; want 1", ans[2].Y)
	}    
    if ans[3].X != 1 {
		t.Errorf("ParseToCoord()= %d; want 1", ans[3].X)
	}    
    if ans[3].Y != 2 {
		t.Errorf("ParseToCoord()= %d; want 2", ans[3].Y)
	}    
}

func TestParseToCoord(t *testing.T) {
	var coords = "(4,3)";

	ans := ParseToCoord(coords)

    if len(ans) != 2 {
		t.Errorf("ParseToCoord()= %d; want 2", len(ans))
	}
    if ans[0] != 4 {
		t.Errorf("ParseToCoord()= %d; want 4", ans[0])
	}   
	if ans[1] != 3 {
		t.Errorf("ParseToCoord()= %d; want 3", ans[1])
	}
}
func TestDetermineMinimumAB(t *testing.T) {
	ans := DetermineMinimum(pA, pB)

	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}

func TestDetermineMinimumAC(t *testing.T) {
	ans := DetermineMinimum(pA, pC)
	if math.Round(ans) != 4.0 {
		t.Errorf("DetermineMinimum()= %f; want 4.242641", ans)
	}
}

func TestDetermineMinimumOfArray1(t *testing.T) {
	points := []Point{pA, pC, pD}

	ans := DetermineMinimumOfArray(points)

	if ans != 3.605551275463989 {
		t.Errorf("DetermineMinimum()= %f; want 3.605551275463989", ans)
	}
}
func TestDetermineMinimumOfArray2(t *testing.T) {
	points := []Point{pA, pB, pC, pD}

	ans := DetermineMinimumOfArray(points)
	
	if ans != 2 {
		t.Errorf("DetermineMinimum()= %f; want 2", ans)
	}
}
