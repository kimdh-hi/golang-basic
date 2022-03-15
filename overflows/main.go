package main

import (
	"fmt"
	"math"
)

func main() {
	var n uint8 = 255
	n++ // 오버플로우 발생 ( 255+1 -> 0) (런타임에 발생)
	fmt.Println(n)

	n2 := uint8(math.MaxUint8 + 1) // 컴파일 타임에 오버플로우 에러 발생
	fmt.Println(n2)

}
