package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		if v := z - (z*z-x)/(2*z); v == z {
			fmt.Printf("%g loops is done", i+1)
			return v
		} else {
			z = v
		}
	}
	return z
}

func main() {
	fmt.Printf("original : %g\n", Sqrt(200))
	fmt.Printf("math     : %g", math.Sqrt(200))
}
