package main

import (
	"errors"
	"fmt"
	"math"
)

// Point holds coods
type Point struct {
	X int
	Y int
}

func main() {
	fmt.Printf("Start App...\n")

	pA := Point{X: 1, Y: 1}
	pB := Point{X: 1, Y: 3}
	pC := Point{X: 4, Y: 4}
	pD := Point{X: 6, Y: 1}

	points := []Point{pA, pB, pC, pD}
	fmt.Println("Minimum distance of all points = ", DetermineMinimumOfArray(points))
}

// DetermineMinimum Determine the minimum of 2 point
func DetermineMinimum(P1 Point, P2 Point) float64 {
	return math.Sqrt(math.Pow(float64(P2.X-P1.X), 2) + math.Pow(float64(P2.Y-P1.Y), 2))
}

// DetermineMinimumOfArray Determine the minimum of a point array
func DetermineMinimumOfArray(points []Point) float64 {
	if len(points) < 1 {
		errors.New("too less points")
	}
	var min float64 = 99999999
	for i := 0; i < len(points); i++ {
		var startPoint = points[i]
		//fmt.Println(startPoint.)

		if startPoint == (Point{}) {
			errors.New("is null")
		}

		min = Calculate(points, startPoint, min, i)
	}

	return min
}

// Calculate min
func Calculate(points []Point, startPoint Point, min float64, startVal int) float64 {
	for i := startVal + 1; i < len(points); i++ {
		var minTemp = DetermineMinimum(startPoint, points[i])
		if minTemp < min {
			min = minTemp
		}
	}

	return min
}
