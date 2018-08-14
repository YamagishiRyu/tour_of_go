package main

import (
	"fmt"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	//	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
	var s string
	for index, a := range i {
		s += fmt.Sprintf("%v", a)
		if index != len(i)-1 {
			s += fmt.Sprintf(".")
		}
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
