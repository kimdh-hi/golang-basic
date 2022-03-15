package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> float
	s1 := "1.123"
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Printf("f1 = %f\n", f1)

	// int -> ascii
	var a1 = strconv.Itoa(1)
	fmt.Printf("a1 = %v\n", a1)

	// ascii -> int
	i, err := strconv.Atoi("-50")
	fmt.Printf("i = %v\n", i)
}
