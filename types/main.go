package main

import "fmt"

func main() {

	// int8: -128 ~ 127
	var i1 int8 = 127
	var i2 int8 = -128
	fmt.Printf("int8 i1=%d, i2=%d\n", i1, i2)

	// uint8 0~255
	var ui1 uint8 = 255
	fmt.Printf("uint9 ui1=%d\n", ui1)

	// rune : unicode
	var rune1 rune = 'a'
	var rune2 rune = '김'
	fmt.Printf("rune1=%v, rune2=%v\n", rune1, rune2)
	fmt.Printf("rune1=%x, rune2=%x\n", rune1, rune2)

	// bool : true, false
	var bool1 bool = true
	var bool2 bool = false
	fmt.Printf("bool1=%v, bool2=%v\n", bool1, bool2)

	// string
	var str string = "hello"
	fmt.Printf("str=%s\n", str)

	// array :  [size]type
	var nums = [3]int{1, 2, 3}
	fmt.Printf("nums type=%T\n", nums)

	// slice
	var names = []string{"kim", "lee", "park"}
	fmt.Printf("names type=%T\n", names)

	// map (key-value)
	var members = map[string]uint8{
		"kim":  27,
		"lee":  25,
		"park": 30,
	}
	fmt.Printf("members type=%T\n", members)

	// struct
	type Member struct {
		name string
		age  uint8
	}
	var member Member
	fmt.Printf("Member type=%T\n", member)

	// pointer
	var x int = 1
	xptr := &x
	fmt.Printf("xptr type=%T, xptr value=%v", xptr, xptr)

}
