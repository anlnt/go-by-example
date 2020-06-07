package main

import(
	"fmt"
)

func main() {
	x := 7
	if x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}

	y := 8
	if y%4 == 0 {
		fmt.Printf("%d is divisible by 4\n", y)
	}

	if num := 18; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
