package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello world")
	distance := 60.8
	fmt.Printf("distance in miles = %f \n", distance*0.62137)

	var a int = 1
	var b float64 = 1.5
	const c int = 10

	fmt.Println(a, b, c)
}
