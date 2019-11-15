package main

import (
	"fmt"
	"math"

	"go-hexagonal-architecture/pkg/random"
)

func main() {
	fmt.Println(4, math.Ceil(math.Log2(4)), math.Log2(4))
	fmt.Println(31, math.Ceil(math.Log2(31)), math.Log2(31))
	fmt.Println(32, math.Ceil(math.Log2(32)), math.Log2(32))
	fmt.Println(33, math.Ceil(math.Log2(33)), math.Log2(33))
	fmt.Println(63, math.Ceil(math.Log2(63)), math.Log2(63))
	fmt.Println(64, math.Ceil(math.Log2(64)), math.Log2(64))
	fmt.Println(65, math.Ceil(math.Log2(65)), math.Log2(65))

	for i := 0; i < 10; i++ {
		fmt.Println(random.String(100))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(random.String(100, random.Alphanumeric, random.Symbols))
	}
}
