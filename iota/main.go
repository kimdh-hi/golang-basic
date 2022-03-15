package main

import "fmt"

/*
iota default 0,1,2 ...
*/

func main() {

	const (
		c1 = iota // default=0
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		cc1 = iota
		cc2
		cc3
	)
	fmt.Println(cc1, cc2, cc3)

	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n3)

	const (
		nn1 = iota * 2 // 0, 2, 4, 6 ...
		nn2
		nn3
		nn4
	)
	fmt.Printf("nn2=%d, nn4=%d\n", nn2, nn4)

	const (
		nnn1 = iota*2 + 1
		nnn2
		nnn3
		nnn4
	)
	fmt.Printf("nnn1=%d, nnn2=%d, nnn4=%d\n", nnn1, nnn2, nnn4)

	// x=-3, y=-5, z=-6
	const (
		x = -(iota + 3) // -3
		_               // -4 (skip)
		y               // -5
		z               // -6
	)
	fmt.Printf("x=%d, y=%d, z=%d", x, y, z)
}
