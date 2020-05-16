package main

import (
    "fmt"
	"errors"
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

func main() {
	fmt.Printf("Start App...\n")
	
	pA := Point {
        X: 1,
        Y: 1,
	}	
	pB := Point {
        X: 1,
        Y: 3,
	}
	pC := Point {
        X: 4,
        Y: 4,
	}	
	pD := Point {
        X: 6,
        Y: 1,
    }

	fmt.Println("DetermineMinimum(pA, pB)", DetermineMinimum(pA, pB))
	fmt.Println("DetermineMinimum(pC, pD)", DetermineMinimum(pC, pD))

	//Point[] points = {pA, pB, pC, pD};
	fmt.Println("Minimum distance of all points = ", DetermineMinimumOfArray(pA, pB, pC, pD))

	if 0 == 0 {
		fmt.Println("0 == 0")
	}

    for i := 0; i <= 4; i++ {
        fmt.Println(i)
	}	
	
	var a [5]int
	fmt.Println("emp:", a)
	fmt.Println("len:", len(a))
	
	if 42 == 42 {
		errors.New("can't work with 42") // ?
    }	
}
