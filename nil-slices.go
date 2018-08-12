package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
}
