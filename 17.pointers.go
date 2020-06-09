package main

import (
	"fmt"
)

func zeroval(i int) {
	i = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

/*
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc0000b4008
*/
