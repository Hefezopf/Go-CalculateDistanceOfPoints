package main

import (
  "testing"
  "calc/point"
)
func TestCalc(t *testing.T) {
    pA := point.Point {
        X: 1,
        Y: 1,
    }	
    pB := point.Point {
        X: 1,
        Y: 3,
    }
    ans := point.DetermineMinimum(pA, pB)
    if ans != 1 {
		t.Errorf("DetermineMinimum()= %d; want 1", ans)
    }
}		
