package main

import (
    "fmt"
	"math"
	"errors"
)

func main() {
	fmt.Printf("App start..\n")

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