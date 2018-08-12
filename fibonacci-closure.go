package main

import (
	"fmt"
)

func fibonacci() func() int {
	a1 := 0
	a2 := 1
	return func() int {
		sum := a1 + a2
		a0 := a1
		a1 = a2
		a2 = sum
		return a0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
