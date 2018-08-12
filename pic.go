package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i, _ := range r {
		v := make([]uint8, dx)
		for j, _ := range v {
			v[j] = uint8(i ^ j)
		}
		r[i] = v
	}
	return r
}

func main() {
	pic.Show(Pic)
}
