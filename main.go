package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"os"
)

// Point holds coods
type Point struct {
	X int
	Y int
}

// E.g.: go run main.go "(4,3) (5,2) (1,1) (1,2)"
func main() {
	fmt.Printf("Start App...\n")

	var points = []Point{}
	if (len(os.Args) > 1) {
		points = ParseAllCoords(os.Args[1]);
	} else {
		var pointsFixed = []Point{}
		pointsFixed = append(pointsFixed, Point{X: 1, Y: 1})
		pointsFixed = append(pointsFixed, Point{X: 1, Y: 3})
		pointsFixed = append(pointsFixed, Point{X: 4, Y: 4})
		pointsFixed = append(pointsFixed, Point{X: 6, Y: 1})
		points = pointsFixed;
	}

	fmt.Println("Minimum distance of all points = ", DetermineMinimumOfArray(points))
	fmt.Println("Maximum distance of all points = ", DetermineMaximumOfArray(points))
}

// ParseAllCoords as string into array of Points. E.g.: "(4,3) (1,2)"
func ParseAllCoords(coords string) []Point {
	var points = []Point{}
	var coordsArray = strings.Split(coords, " ")
	for i := 0; i < len(coordsArray); i++ {
		xy := ParseToCoord(coordsArray[i])
		points = append(points, Point{X: xy[0], Y: xy[1]})
	}

	return points
}

// ParseToCoord as string into array of int. E.g.: (4,3)
func ParseToCoord(coords string) []int {
	var res = []int{0, 0}

	x, err := strconv.Atoi(coords[1:2])
	y, err := strconv.Atoi(coords[3:4])
	res[0] = x
	res[1] = y
	if err != nil {
		fmt.Print("err:", err)
	}

	return res
}

// DetermineDistance Determine the distance of 2 points
func DetermineDistance(P1 Point, P2 Point) float64 {
	return math.Sqrt(math.Pow(float64(P2.X-P1.X), 2) + math.Pow(float64(P2.Y-P1.Y), 2))
}

// DetermineMaximumOfArray Determine the maximum of a point array
func DetermineMaximumOfArray(points []Point) float64 {
	if len(points) < 1 {
		errors.New("too less points")
	}
	var max float64 = 0
	for i := 0; i < len(points); i++ {
		var startPoint = points[i]
		if startPoint == (Point{}) {
			errors.New("startPoint is null")
		}

		max = calculateMaximum(points, startPoint, max, i)
	}

	return max
}

// DetermineMinimumOfArray Determine the minimum of a point array
func DetermineMinimumOfArray(points []Point) float64 {
	if len(points) < 1 {
		errors.New("too less points")
	}
	var min float64 = math.MaxFloat64
	for i := 0; i < len(points); i++ {
		var startPoint = points[i]
		if startPoint == (Point{}) {
			errors.New("startPoint is null")
		}

		min = calculateMinimum(points, startPoint, min, i)
	}

	return min
}

// Calculate maximum
func calculateMaximum(points []Point, startPoint Point, max float64, startVal int) float64 {
	for i := startVal + 1; i < len(points); i++ {
		var maxTemp = DetermineDistance(startPoint, points[i])
		if maxTemp > max {
			max = maxTemp
		}
	}

	return max
}

// Calculate minimum
func calculateMinimum(points []Point, startPoint Point, min float64, startVal int) float64 {
	for i := startVal + 1; i < len(points); i++ {
		var minTemp = DetermineDistance(startPoint, points[i])
		if minTemp < min {
			min = minTemp
		}
	}

	return min
}
