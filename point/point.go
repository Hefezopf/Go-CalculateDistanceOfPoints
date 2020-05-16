package point

import (
	"math"
)

type Point struct {
	X int
	Y int
}

func DetermineMinimum(P1 Point, P2 Point) float64 {  
	return math.Sqrt(math.Pow(float64(P2.X-P1.X),2) + math.Pow(float64(P2.Y-P1.Y),2) );
}

func DetermineMinimumOfArray(P1 Point, P2 Point, P3 Point, P4 Point) float64 {  
	return math.Sqrt(math.Pow(float64(P2.X-P1.X),2) + math.Pow(float64(P2.Y-P1.Y),2) );
}


