package main

import (
    "fmt"
	"errors"
	"calc/point"
)

func main() {
	fmt.Printf("Start App...\n")
	
	pA := point.Point {
        X: 1,
        Y: 1,
	}	
	pB := point.Point {
        X: 1,
        Y: 3,
	}
	pC := point.Point {
        X: 4,
        Y: 4,
	}	
	pD := point.Point {
        X: 6,
        Y: 1,
    }

	fmt.Println("DetermineMinimum(pA, pB)", point.DetermineMinimum(pA, pB))
	fmt.Println("DetermineMinimum(pC, pD)", point.DetermineMinimum(pC, pD))

	//Point[] points = {pA, pB, pC, pD};
	fmt.Println("Minimum distance of all points = ", point.DetermineMinimumOfArray(pA, pB, pC, pD))

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
