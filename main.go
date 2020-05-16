package main
import (
    "fmt"
	"math"
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

	fmt.Println("DetermineMinimum2()", point.DetermineMinimum(pA, pB))
	fmt.Println("DetermineMinimum2()", point.DetermineMinimum(pC, pD))
	//fmt.Println("determine()", determine())

	
	//Point[] points = {pA, pB, pC, pD};
	
	//System.out.println("Distance A, C = "+ Point.determineMinimum(pA, pC));
	
	//double minDistanceOfArray = Point.determineMinimumOfArray(points);
	//System.out.println("Minimum distance of all points = "+ minDistanceOfArray);


	fmt.Println(math.Sin(1))

	if 0 == 0 {
		fmt.Println("0 == 0")
	}

    for i := 0; i <= 4; i++ {
        fmt.Println(i)
	}	
	
	var a [5]int
	fmt.Println("emp:", a)
	fmt.Println("len:", len(a))
		
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	
	if 42 == 42 {
		errors.New("can't work with 42") // ?
    }	
}

func plus(a int, b int) int {
	return a + b
}


func calc() int {
    return 1
}
