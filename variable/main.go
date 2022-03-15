package main

import "fmt"

func main() {

	var age int = 27
	fmt.Println("age = ", age)

	var name = "kim" // 사용하지 않는 경우 컴파일 에러
	_ = name         // 사용하지 않는 변수에 대한 컴파일 에러 해결

	language := "golang"
	fmt.Println(language)

	// 변수 여러 개 선언
	a, b := 1, "str"
	fmt.Println(a, b)

	// swap
	i, j := 1, 2
	i, j = j, i
	fmt.Println("i=", i, "j=", j)

	// 캐스팅
	var num1 = 5
	var num2 = 5.2
	num1 = int(num2)
	fmt.Println("num1=", num1)

}
