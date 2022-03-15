package main

import "fmt"

func main() {
	num := 1
	str := "str"
	flag := true

	// %v : 포멧팅 모든 타입 OK
	fmt.Printf("num = %v, str = %v, flag = %v\n", num, str, flag)
	fmt.Printf("num = %v\n str = %v\n", num, str)

	// %T : 데이터 타입 출력
	fmt.Printf("num type: %T, str type: %T, flag type: %T\n", num, str, flag)

	// %.3f : 소수점 포멧팅
	fnum := 1.12345678
	fmt.Printf("fnum = %.5f\n", fnum)

}

/*
%d -> decimal
%f -> float
%s -> string
%q -> double-quoted string
%v -> value (any)
%#v -> a Go-syntax representation of the value
%T -> value Type
%t -> bool (true or false)
%p -> pointer (address in base 16, with leading 0x)
%c -> char (rune) represented by the corresponding Unicode code poi
*/
