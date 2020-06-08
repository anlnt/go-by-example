package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Printf("1 + 2 = %d\n", res)

	res = plusPlus(1, 2, 3)
	fmt.Printf("1 + 2 + 3 = %d\n", res)
}
