package main

import (
	"fmt"
)

func gcdr(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcdr(b, a%b)
	}
}

func gcd(a, b int) int {
	if a < b {
		return gcdr(b, a)
	} else {
		return gcdr(a, b)
	}
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%d\n", gcd(a, b))
}
