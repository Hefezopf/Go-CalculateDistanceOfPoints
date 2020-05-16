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
/*
public static double determineMinimumOfArray(Point[] points) {	
	if(points.length < 1) {
		throw new IllegalArgumentException("too less points");
	}
	double min = Double.POSITIVE_INFINITY;
	for (int i = 0; i < points.length; i++) {
		Point startPoint = points[i];
		if(startPoint==null) {
			throw new IllegalArgumentException("");
		}
		min = calculate(points, startPoint, min, i);
	}
	
	return min;
}

private static double calculate(Point[] points, Point startPoint, double min, int startVal) {
	for (int i = startVal+1; i < points.length; i++) {
		double minTemp = determineMinimum(startPoint, points[i]);
		if(minTemp < min) {
			min = minTemp;
		}
	}
	return min;
}
*/


